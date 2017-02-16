//go:generate swagger generate spec
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/toshim45/jajak/domains/poll"
	"github.com/toshim45/jajak/handlers"
	"github.com/toshim45/jajak/utils"
	"gopkg.in/mgo.v2"
)

const (
	mongoURI = "mongodb://localhost:27017/jajak"
	dbName   = "jajak"
	port     = "8071"
)

func main() {
	utils.CommonPanicHandler()
	createRoutes()
}

func createRoutes() {
	session := initMongo()
	db := session.DB(dbName)

	pollService := poll.NewService(db)

	pingHandler := handlers.NewPingHandler(session)
	pollHandler := handlers.NewPollHandler(pollService)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/polls", pollHandler.GetPolls).Methods("GET")

	chainHandler := alice.New(utils.LoggingHandler)

	log.Printf("server up at port %s", port)
	http.ListenAndServe(":"+port, chainHandler.Then(r))
	defer session.Close()
}

func initMongo() *mgo.Session {
	session, err := mgo.Dial(mongoURI)
	utils.ThrowPanic(err)
	log.Printf("connected to mongo on %s", mongoURI)
	return session
}

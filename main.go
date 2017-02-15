package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	createRoutes()
}

func createRoutes() {
	session := initMongo()
	db := session.DB(dbName)

	r := mux.NewRouter()

	pingHandler := handlers.NewPingHandler(session)
	pollHandler := handlers.NewPollHandler(db)

	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/polls", pollHandler.GetPolls).Methods("GET")

	log.Printf("server up at port %s", port)
	http.ListenAndServe(":"+port, r)
	defer session.Close()
}

func initMongo() *mgo.Session {
	session, err := mgo.Dial(mongoURI)
	utils.ThrowPanic(err)
	log.Printf("connected to mongo on %s", mongoURI)
	return session
}

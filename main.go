//go:generate swagger generate spec
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/toshim45/jajak/domains/poll"
	"github.com/toshim45/jajak/handlers"
	"github.com/toshim45/jajak/utils"
	"gopkg.in/mgo.v2"
)

func main() {
	utils.CommonPanicHandler()
	config := utils.GetConfig()
	createRoutes(config)
}

func createRoutes(config utils.Config) {
	session := initMongo(config)
	db := session.DB(config.MongoDBName)

	pollService := poll.NewService(db)

	pingHandler := handlers.NewPingHandler(session)
	pollHandler := handlers.NewPollHandler(pollService)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/polls", pollHandler.GetPolls).Methods("GET")

	chainHandler := alice.New(utils.LoggingHandler)

	if config.EnableSwagger {
		log.Printf("swagger enabled, loading CORS with origin: %s", config.AllowedOrigin)
		chainHandler = chainHandler.Append(utils.EnableCors(config).Handler)
	}

	log.Printf("server up at port %s", config.Port)
	http.ListenAndServe(":"+config.Port, chainHandler.Then(r))
	defer session.Close()
}

func initMongo(c utils.Config) *mgo.Session {
	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", c.MongoHost, c.MongoPort, c.MongoDBName)
	session, err := mgo.Dial(mongoURI)
	utils.ThrowPanic(err)
	log.Printf("connected to mongo on %s", mongoURI)
	return session
}

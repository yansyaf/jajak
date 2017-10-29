package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/toshim45/jajak/config"
	"github.com/toshim45/jajak/httphandler"
	"github.com/toshim45/jajak/httputil"
	"github.com/toshim45/jajak/poll"
	"github.com/toshim45/jajak/uptime"

	"gopkg.in/mgo.v2"
)

func main() {
	httputil.CommonPanicHandler()
	envConfig := config.NewEnv()
	createRoutes(envConfig)
}

func createRoutes(envConfig config.Environment) {
	session := initMongo(envConfig)
	db := session.DB(envConfig.MongoDBName)
	upTime := uptime.New()

	pollService := poll.NewService(db)

	pingHandler := httphandler.NewPingHandler(session, upTime)
	pollHandler := httphandler.NewPollHandler(pollService)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/polls", pollHandler.GetPolls).Methods("GET")
	r.HandleFunc("/polls/{id}", pollHandler.GetPollById).Methods("GET")

	chainHandler := alice.New(httputil.LoggingHandler)

	if envConfig.EnableSwagger {
		log.Printf("swagger enabled, loading CORS with origin: %s", envConfig.AllowedOrigin)
		chainHandler = chainHandler.Append(httputil.EnableCors(envConfig).Handler)
	}

	log.Printf("server up at port %s", envConfig.Port)
	http.ListenAndServe(":"+envConfig.Port, chainHandler.Then(r))
	defer session.Close()
}

func initMongo(c config.Environment) *mgo.Session {
	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", c.MongoHost, c.MongoPort, c.MongoDBName)
	session, err := mgo.Dial(mongoURI)
	httputil.ThrowPanic(err)
	log.Printf("connected to mongo on %s", mongoURI)
	return session
}

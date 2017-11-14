package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/toshim45/jajak/config"
	"github.com/toshim45/jajak/httphandler"
	"github.com/toshim45/jajak/httputil"
	"github.com/toshim45/jajak/survey"
	"github.com/toshim45/jajak/uptime"

	"gopkg.in/mgo.v2"
)

func main() {
	httputil.CommonPanicHandler()
	envConfig := config.NewEnv()
	mgoSession := initMongo(envConfig)
	router := createRoutes(envConfig, mgoSession)
	chainHandler := alice.New(httputil.LoggingHandler)

	if envConfig.EnableSwagger {
		log.Printf("swagger enabled, loading CORS with origin: %s", envConfig.AllowedOrigin)
		chainHandler = chainHandler.Append(httputil.EnableCors(envConfig).Handler)
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{Addr: ":" + envConfig.Port, Handler: chainHandler.Then(router)}

	go listenToSigTerm(stopChan, server, mgoSession)

	log.Printf("server up at port %s", envConfig.Port)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalln("http closed with Error:", err)
		}
	}
}

func listenToSigTerm(stopChan chan os.Signal, server *http.Server, mgoSession *mgo.Session) {
	<-stopChan

	log.Println("Shutting down server in ", httphandler.ShuttingDownDelay.String())
	httphandler.IsShuttingDown = true

	time.Sleep(httphandler.ShuttingDownDelay)
	mgoSession.Close()

	time.Sleep(httphandler.ShuttingDownDelay)
	ctx, cancel := context.WithTimeout(context.Background(), httphandler.ShuttingDownDelay)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown: %v", err)
	}
	log.Println("Bye")
}

func createRoutes(envConfig config.Environment, session *mgo.Session) *mux.Router {
	db := session.DB(envConfig.MongoDBName)
	upTime := uptime.New()

	surveyService := survey.NewMongoService(db)

	pingHandler := httphandler.NewPing(upTime, func() error { return session.Ping() })
	//	pingHandler := httphandler.NewPing(upTime)
	surveyHandler := httphandler.NewSurvey(surveyService)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/surveys", surveyHandler.GetSurveys).Methods("GET")
	r.HandleFunc("/surveys", surveyHandler.StoreSurvey).Methods("POST")
	r.HandleFunc("/surveys/{id}", surveyHandler.GetSurveyById).Methods("GET")
	r.HandleFunc("/surveys/{id}/polls", surveyHandler.StorePoll).Methods("POST")

	return r
}

func initMongo(c config.Environment) *mgo.Session {
	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", c.MongoHost, c.MongoPort, c.MongoDBName)
	session, err := mgo.Dial(mongoURI)
	httputil.ThrowPanic(err)
	log.Printf("connected to mongo on %s", mongoURI)
	return session
}

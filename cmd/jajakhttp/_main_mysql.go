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

	"github.com/toshim45/jajak/config"
	"github.com/toshim45/jajak/httphandler"
	"github.com/toshim45/jajak/httputil"
	"github.com/toshim45/jajak/survey"
	"github.com/toshim45/jajak/uptime"

	//	"gopkg.in/mgo.v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
)

func main() {
	httputil.CommonPanicHandler()
	envConfig := config.NewEnv()
	//	mgoSession := initMongo(envConfig)
	mysqlDB := initMySQL(envConfig)
	//	router := createRoutes(envConfig, mgoSession)
	router := createRoutes(envConfig, mysqlDB)
	chainHandler := alice.New(httputil.LoggingHandler)

	if envConfig.EnableSwagger {
		log.Printf("swagger enabled, loading CORS with origin: %s", envConfig.AllowedOrigin)
		chainHandler = chainHandler.Append(httputil.EnableCors(envConfig).Handler)
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{Addr: ":" + envConfig.Port, Handler: chainHandler.Then(router)}

	//	go listenToSigTerm(stopChan, server, mgoSession)
	go listenToSigTerm(stopChan, server, mysqlDB)

	log.Printf("server up at port %s", envConfig.Port)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalln("http closed with Error:", err)
		}
	}
}

//func listenToSigTerm(stopChan chan os.Signal, server *http.Server, mgoSession *mgo.Session) {
func listenToSigTerm(stopChan chan os.Signal, server *http.Server, db *sqlx.DB) {
	<-stopChan

	log.Println("Shutting down server in ", httphandler.ShuttingDownDelay.String())
	httphandler.IsShuttingDown = true

	time.Sleep(httphandler.ShuttingDownDelay)
	//	mgoSession.Close()
	db.Close()

	time.Sleep(httphandler.ShuttingDownDelay)
	ctx, cancel := context.WithTimeout(context.Background(), httphandler.ShuttingDownDelay)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown: %v", err)
	}
}

//func createRoutes(envConfig config.Environment, session *mgo.Session) *mux.Router {
//	db := session.DB(envConfig.MongoDBName)
func createRoutes(envConfig config.Environment, db *sqlx.DB) *mux.Router {
	upTime := uptime.New()

	surveyService := survey.New(db)

	//	pingHandler := httphandler.NewPing(session, upTime)
	pingHandler := httphandler.NewPing(db, upTime)
	surveyHandler := httphandler.NewSurvey(surveyService)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.GetPing).Methods("GET")
	r.HandleFunc("/surveys", surveyHandler.GetSurveys).Methods("GET")
	r.HandleFunc("/surveys", surveyHandler.StoreSurvey).Methods("POST")
	r.HandleFunc("/surveys/{id}", surveyHandler.GetSurveyById).Methods("GET")
	r.HandleFunc("/surveys/{id}/polls", surveyHandler.StorePoll).Methods("POST")

	return r
}

//func initMongo(c config.Environment) *mgo.Session {
//	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", c.MongoHost, c.MongoPort, c.MongoDBName)
//	session, err := mgo.Dial(mongoURI)
//	httputil.ThrowPanic(err)
//	log.Printf("connected to mongo on %s", mongoURI)
//	return session
//}

func initMySQL(c config.Environment) *sqlx.DB {
	log.Printf("Trying to connect to %s db..", c.MySQLDBHost)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", c.MySQLDBUser, c.MySQLDBPassword, c.MySQLDBHost, c.MySQLDBName)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(c.MySQLDBConnLimit)
	db.SetMaxIdleConns(c.MySQLDBConnLimit)

	log.Printf("Connected to db: %s@%s/%s", c.MySQLDBUser, c.MySQLDBHost, c.MySQLDBName)
	return db
}

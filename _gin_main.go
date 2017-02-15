package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
)

const (
	MongoURI = "mongodb://localhost:27017/jajak"
	DbName   = "jajak"
)

func PanicError(e error) {
	if e != nil {
		panic(e)
	}
}

func ContextError(c *gin.Context, e error) {
	if e != nil {
		c.Error(e)
		PanicError(e)
	}
}

func init() {
	session, err := mgo.Dial(MongoURI)
	PanicError(err)

	Session = session
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/polls", listPoll)
	router.POST("/polls", submitPoll)

	router.Run(":8071")

	defer Session.Close()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func listPoll(c *gin.Context) {
	c.JSON(200, ListPoll(c, Session.DB(DbName)))
}

func submitPoll(c *gin.Context) {
	SubmitPoll(c, Session.DB(DbName))
	c.Writer.WriteHeader(201)
}

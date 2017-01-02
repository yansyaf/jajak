package main

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

type Poll struct {
	title   string `bson:"title" json:"title"`
	creator string `bson:"creator" json:"creator"`
	//	items   []string
}

const (
	CollName = "poll3"
)

func prints(polls *[]Poll) {
	for _, poll := range *polls {
		print(&poll)
	}
}

func print(poll *Poll) {
	fmt.Println("[poll]> " + poll.title + " " + poll.creator)
}

func ListPoll(context *gin.Context, db *mgo.Database) []Poll {
	//	db := context.MustGet("db").(*mgo.Database)
	polls := []Poll{}
	err := db.C(CollName).Find(nil).All(&polls)
	ContextError(context, err)

	prints(&polls)

	return polls
}

func SubmitPoll(context *gin.Context, db *mgo.Database) {
	poll := Poll{}
	err := context.BindJSON(&poll)
	ContextError(context, err)

	print(&poll)
	err = db.C(CollName).Insert(poll)
	ContextError(context, err)
}

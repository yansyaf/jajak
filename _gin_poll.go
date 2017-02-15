package main

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

type Poll struct {
	Title   string   `bson:"title" json:"title"`
	Creator string   `bson:"creator" json:"creator"`
	Items   []string `bson:"items" json:"items"`
}

const (
	CollName = "polls"
)

func prints(polls *[]Poll) {
	for _, poll := range *polls {
		print(&poll)
	}
}

func print(poll *Poll) {
	fmt.Printf("[poll]> %s %s %d\r\n", poll.Title, poll.Creator, len(poll.Items))
	if len(poll.Items) > 0 {
		fmt.Printf("[poll][item]> ")
		for _, item := range poll.Items {
			fmt.Printf("%s\t", item)
		}
		fmt.Printf("\r\n")
	}
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

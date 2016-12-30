package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type Poll struct {
	Title   string
	Creator string
	//	Items   []string
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/polls", listPoll)

	router.Run(":8071")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func listPoll(c *gin.Context) {
	//	items := new []string{"mie","nasi","baso"}
	result := Poll{Title: "food", Creator: "awibowo@cacucu.com"}
	c.JSON(200, result)
}

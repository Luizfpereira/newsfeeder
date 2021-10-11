package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run("127.0.0.1:8080")

	// Parei no minuto 34
	//https://www.youtube.com/watch?v=LOn1GUsjOF4&ab_channel=DavidAlsh

}

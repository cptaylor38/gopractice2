package main

import "github.com/gin-gonic/gin"

func httpMain() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()

	r.GET("/test/:name", func(c *gin.Context) {

	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", IndexHandler)
	http.ListenAndServe(":8080", r)
}

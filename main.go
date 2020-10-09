package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		fmt.Println("Hello World")
		c.String(http.StatusOK, "Hello World")
	})

	router.Run()
}

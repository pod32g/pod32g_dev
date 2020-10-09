package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pod32g/pod32g_com/Controllers"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Hello World")
		c.String(http.StatusOK, "Hello World")
	})

	authentication := Controllers.Authentication{
		Router: router,
	}

	authentication.Init()

	router.Run()

}

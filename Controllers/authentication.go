package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authentication struct {
	Router *gin.Engine
}

func (a *Authentication) Authenticate(c *gin.Context) {

}

func (a *Authentication) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (a *Authentication) Init() {
	auth := a.Router.Group("/auth")
	{
		auth.GET("/", a.test)
	}
}

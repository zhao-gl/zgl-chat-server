package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Common() {
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Gin services!")
	})
	Router.GET("/404", func(c *gin.Context) {
		c.String(http.StatusOK, "404!")
	})
}

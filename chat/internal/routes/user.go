package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var userPath = "/user"

func User() {
	Router.GET(userPath, func(c *gin.Context) {
		c.String(http.StatusOK, "Zhao's profile!")
	})
}

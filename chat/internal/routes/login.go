package routes

import (
	"chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

var loginPath = "/login"

func Login() {
	Router.GET(loginPath, func(c *gin.Context) {
		handlers.LoginHandler(c)
	})
	Router.GET(loginPath+"/reg", func(c *gin.Context) {
		handlers.RegHandler(c)
	})
}

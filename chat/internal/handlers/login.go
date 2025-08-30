package handlers

import (
	"chat/internal/models"
	"chat/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello Zhao!")
	username := c.DefaultQuery("username", "none")
	password := c.DefaultQuery("password", "none")
	var params = models.LoginModel{
		Username: username,
		Password: password,
	}
	services.LoginService(c, params)
}

func RegHandler(c *gin.Context) {
	c.String(http.StatusOK, "Reg Zhao!")
	//services.RegService(data)
}

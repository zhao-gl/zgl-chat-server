package handlers

import (
	"chat/internal/models"
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, models.Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

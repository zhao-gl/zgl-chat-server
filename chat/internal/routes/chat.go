package routes

import (
	"chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Chat() {
	// 创建新会话
	Router.POST("/chat/new", func(c *gin.Context) {
		handlers.HandleCreateSession(c)
	})

	// 查询会话
	Router.GET("/chat/querySessionById", func(c *gin.Context) {
		handlers.HandleQuerySessionById(c)
	})

	// 获取历史会话
	Router.GET("/chat/queryAllSessions", func(c *gin.Context) {
		handlers.HandleAllSessions(c)
	})

	// 删除会话
	Router.GET("/chat/delSession", func(c *gin.Context) {
		handlers.HandleDelSession(c)
	})

	// 聊天接口
	Router.POST("/chat/ask", func(c *gin.Context) {
		handlers.HandleAsk(c)
	})
}

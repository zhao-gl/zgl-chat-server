package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Router 创建一个默认的 Gin 路由引擎
var Router = gin.Default()

func OpenServer() {
	Router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		//AllowHeaders:     []string{"Origin"},
		//ExposeHeaders:    []string{"Content-Length"},
		//AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))

	Common()
	User()
	Login()
	Chat()
	// 启动 HTTP 服务并监听 3000 端口
	err := Router.Run(":3000")
	if err != nil {
		panic(err)
	}

}

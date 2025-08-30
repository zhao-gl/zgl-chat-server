package main

import (
	"chat/config"
	"chat/internal/routes"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// 加载 .env.ai 文件
	err := godotenv.Load(".env.ai")
	if err != nil {
		log.Fatal("Error loading .env.ai file")
	}
	// 初始化数据库
	config.InitDB()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			println("数据库关闭失败:", err)
		}
	}(config.DB)

	routes.OpenServer()
}

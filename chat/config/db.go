package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
)

var DB *sql.DB

func InitDB() {
	var err error
	// 数据库连接字符串格式: username:password@tcp(host:port)/database
	dsn := "root:root@tcp(127.0.0.1:3306)/chat"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 测试连接
	if err = DB.Ping(); err != nil {
		log.Fatal("数据库Ping失败:", err)
	}

	fmt.Println("数据库连接成功")
}

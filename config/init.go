package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	// 数据库参数
	DatabaseURI  string
	DatabaseName string
	JwtKey       string
)

func init() {
	// 加载配置文件
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading config.env file", err)
	}

	// 加载配置文件
	DatabaseURI = os.Getenv("DATABASE_URI")
	DatabaseName = os.Getenv("DATABASE_NAME")
	JwtKey = os.Getenv("JWT_KEY")
}

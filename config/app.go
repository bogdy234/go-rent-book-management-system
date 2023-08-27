package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables(path string) {
	err := godotenv.Load(path)

	if err != nil {
		panic(err)
	}
}

func ConnectToDatabase() {
	dsn := os.Getenv("DB_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}

func Init(envPath ...string) {
	p := ".env"
	if len(envPath) > 0 {
		p = envPath[0]
	}

	LoadEnvVariables(p)
	ConnectToDatabase()
}

func GetDB() *gorm.DB {
	return DB
}

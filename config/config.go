package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbURL := os.Getenv("DB_URL")

	db, dbError := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if dbError != nil {
		panic("Failed to connect to Database")
	}

	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill database connection")
	}
	err = dbSQL.Close()
	if err != nil {
		panic("Failed to kill database connection")
	}
}

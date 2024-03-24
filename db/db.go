package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	host = os.Getenv("DB_HOST")
	port,_  = strconv.Atoi(os.Getenv("DB_PORT"))
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbName   = os.Getenv("DB_NAME")
)


func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err!=nil {
		fmt.Println("An error occurred:", err)
		return nil
	}

	return db
}
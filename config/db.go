package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbName   = "test"
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
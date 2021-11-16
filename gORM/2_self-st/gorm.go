package main

import (
	"fmt"
	"log"
  	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

const (
	
	Host = "localhost"
	DUser = "jasurbek"
	Password = "1001"
	DbName = "example"
	Port = "5432"
	SslMode = "disable"
	TimeZone = "Asia/Tashkent"
)

var dsn = fmt.Sprintf(

		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Host, DUser, Password, DbName, Port, SslMode, TimeZone,
	)

func DbConnection () (*gorm.DB) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}

	return db
}	
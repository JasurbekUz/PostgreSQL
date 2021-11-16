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


var user User
var postUser PostUser
var users []User

var dsn = fmt.Sprintf(

		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Host, DUser, Password, DbName, Port, SslMode, TimeZone,
	)

func Database (string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}	

func main () {

	db, err := Database(dsn)

	if err != nil {
		log.Fatalf("db connection error: %v", err)	
	}

	postUser = PostUser {

		FullName: "Ikrom Akromov",
		UserName: "ikrom",
		PhoneNumber: "856542569",
	}

	var p_user = map[string]interface{}{

		"full_name" 	: postUser.FullName,
		"user_name" 	: postUser.UserName,
		"phone_number"  : postUser.PhoneNumber,
	}

	db.Model(&user).Create(p_user)

	db.Last(&user)

	fmt.Println(user)
}
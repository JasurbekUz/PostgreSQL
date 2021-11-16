package main

import (

	"log"
)

var user User



func main () {

	db := DbConnection()

	//db.Model(&User{}).Where("user_name = ?", "vali").Update("phone_number", " --- ")//.Scan(user)	

	//db.Model(&user).Where("active = ?", true).Update("name", "hello")
	//db.Model(&user).Update("name", "hello")

	//db.Where("user_name = ?", "vali").Find(&user)
	//db.Where("name = ?", "jinzhu").Delete(&email)

	db.Where("user_name = ?", "jinzhu").Delete(&User{})

	db.Exec("delete from users where user_name = ?", "ikrom")

	log.Printf()
}
package main

import (
	"time"
)

type User struct {
	UserId uint         
	FullName string
	UserName string
	PhoneNumber string
	CreatedAt time.Time
}

type PostUser struct {
	FullName string
	UserName string
	PhoneNumber string
}

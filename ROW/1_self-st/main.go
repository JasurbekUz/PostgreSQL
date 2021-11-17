package main

import (

	"log"
)

const (
	Post int = iota
	Get 
	GetAll 
	Update 
	Delete

)

var newUser PostUser
var user User
var users []User



func main () {

	comand := Update

	if comand == Get {					//GET

		id := 18
		row := GetUser(id)

		log.Println(row)

	} else if comand == GetAll {		// GETALL

		rows := GetAllUsers()

		log.Println(rows)
	
	} else if comand == Post {			// POST

		newUser = PostUser {
			FullName: "Doniyor Doniyorov",
			UserName: "doni",
			PhoneNumber: "998988787",
		}

		row := PostUsers(newUser)

		log.Println(row)

	} else if comand == Delete {		// DELETE

		id := 2

		row := DelUsers(id)

		log.Println(row)
	} else if comand == Update {		// UPDATE

		log.Println("no any options")
	}
}

func GetUser (id int) User {

	db := DataBase()

	defer db.Close()

	err := db.QueryRow(
			"select * from users where user_id = $1",
			id,
		).Scan(
			&user.UserId,
			&user.FullName,
			&user.UserName,
			&user.PhoneNumber,
			&user.CreatedAt,
		)

	if err != nil {
		
		log.Fatalf("reading rows error: %v", err)
	}

	return user
}

func GetAllUsers () []User {

	db := DataBase()

	rows, err := db.Query("select * from users")	

	if err != nil {
		
		log.Fatalf("reading rows error: %v", err)
	}

	for rows.Next() {

		rows.Scan(
			&user.UserId,
			&user.FullName,
			&user.UserName,
			&user.PhoneNumber,
			&user.CreatedAt,
		)

		users = append(users, user)
	}


	return users
}

func PostUsers (newUser PostUser) User {

	db := DataBase()

	err := db.QueryRow(
		"insert into users (full_name, user_name, phone_number) values ($1, $2, $3) returning * ",
		newUser.FullName,
		newUser.UserName,
		newUser.PhoneNumber,
	).Scan(
		&user.UserId,
		&user.FullName,
		&user.UserName,
		&user.PhoneNumber,
		&user.CreatedAt,
	)


	if err != nil {
		
		log.Fatalf("reading rows error: %v", err)
	}

	return user
}

func DelUsers (id int) User {

	db := DataBase()

	log.Println(id)

	err := db.QueryRow(
		"delete from users where user_id = $1 returning * ",
		id,
	).Scan(
		&user.UserId,
		&user.FullName,
		&user.UserName,
		&user.PhoneNumber,
		&user.CreatedAt,
	)

	if err != nil {
		
		log.Fatalf("no info: %v", err)
	}

	return user
}

/*func DelUsers (id int) User {

	db := DataBase()

	log.Println(id)

	err := db.QueryRow(
		"delete from users where user_id = $1 returning * ",
		id,
	).Scan(
		&user.UserId,
		&user.FullName,
		&user.UserName,
		&user.PhoneNumber,
		&user.CreatedAt,
	)

	if err != nil {
		
		log.Fatalf("no info: %v", err)
	}

	return user
}*/
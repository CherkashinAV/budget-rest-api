package database

import (
	"database/sql"
	"fmt"

)

type User struct {
	id int
	name string
	surname string
	available_funds int
	lang string
	theme string
}

func GetUsers() ([]User, error) {
	var users []User
	var err error
	var rows *sql.Rows
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err = db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("getUsers: %v", err)
	}
	
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.id, &usr.name, &usr.surname, &usr.available_funds, &usr.lang, &usr.theme); err != nil {
			return nil, fmt.Errorf("getUsers: %v", err)
		}
		users = append(users, usr)
	}
	return users, nil
} 
package database

import (
	"database/sql"
	"fmt"

)

type User struct {
	id int64
	name string
	surname string
	available_funds int64
	lang string
	theme string
}

func (db DataBase) getUsers() ([]User, error) {
	var users []User
	var err error
	var rows *sql.Rows
	rows, err = db.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("getUsers: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.id, &usr.name, &usr.surname, &usr.available_funds, &usr.lang, &usr.theme); err != nil {
			return nil, fmt.Errorf("getUsers: %v", err)
		}
		users = append(users, usr)
	}
	return users, nil
} 
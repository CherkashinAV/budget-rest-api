package database

import (
	"database/sql"
	"fmt"
)

func GetUsers() ([]User, error) {
	var users []User

	db, err := Connect()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("getUsers: %v", err)
	}
	defer rows.Close()

	db.Close()

	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Surname, &usr.Available_funds, &usr.Lang, &usr.Theme); err != nil {
			return nil, fmt.Errorf("getUsers: %v", err)
		}
		users = append(users, usr)
	}
	return users, nil
} 


func GetUserById(id string) (User, error) {
	var usr User
	db, err := Connect()
	if err != nil {
		return usr, err
	}
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	db.Close()
	if err := row.Scan(&usr.ID, &usr.Name, &usr.Surname, &usr.Available_funds, &usr.Lang, &usr.Theme); err != nil {
		if err == sql.ErrNoRows {
         return usr, fmt.Errorf("userById %s: no such user", id)
      }
      return usr, fmt.Errorf("albumsById %s: %v", id, err)
	}
	return usr, nil
}


func GetAllUserTransactions(user string) ([]Transaction, error) {
	var transactions []Transaction

	db, err := Connect()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM transactions WHERE user = ?", user)
	if err != nil {
		return nil, fmt.Errorf("getAllUserTransactions: %v", err)
	}
	defer rows.Close()

	db.Close()
	for rows.Next() {
		var tr Transaction
		if err := rows.Scan(&tr.ID, &tr.User, &tr.Date, &tr.Ante, &tr.Desc); err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}
	return transactions, nil
}


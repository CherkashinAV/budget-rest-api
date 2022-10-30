package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DataBase struct {
	db *sql.DB
}

func New(db *sql.DB) *DataBase {
	return &DataBase{
		db: db,
	}
}

func Connect() (*DataBase, error){
	cfg := SQLConfig()
	var err error
	var db *sql.DB
   db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return New(db), err
}

func SQLConfig() mysql.Config{
	cfg := mysql.Config {
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "budget",
    }
	return cfg
}


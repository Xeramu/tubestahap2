package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	dsn := "root:root@tcp(mysql:3306)/tubesdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	createTable()
	fmt.Println("USER DB CONNECTED")
}

func createTable() {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100),
		password VARCHAR(255),
		role VARCHAR(50),
		alamat TEXT,
		preferensi TEXT
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
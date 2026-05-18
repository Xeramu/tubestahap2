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

	fmt.Println("ORDER DB CONNECTED")
}

func createTable() {

	query := `
	CREATE TABLE IF NOT EXISTS orders (
		order_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT,
		resi VARCHAR(100),
		nama_barang VARCHAR(100),
		berat INT,
		dimensi VARCHAR(100),
		jenis VARCHAR(100),
		alamat_pengirim TEXT,
		alamat_penerima TEXT,
		status VARCHAR(50),
		eta VARCHAR(100)
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
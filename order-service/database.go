// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func ConnectDB() {

// 	dsn := "root:root@tcp(mysql:3306)/tubesdb"

// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	DB = db

// 	createTable()

// 	fmt.Println("ORDER DB CONNECTED")
// }

// func createTable() {

// 	query := `
// 	CREATE TABLE IF NOT EXISTS orders (
// 		order_id INT AUTO_INCREMENT PRIMARY KEY,
// 		user_id INT,
// 		resi VARCHAR(100),
// 		nama_barang VARCHAR(100),
// 		berat INT,
// 		dimensi VARCHAR(100),
// 		jenis VARCHAR(100),
// 		alamat_pengirim TEXT,
// 		alamat_penerima TEXT,
// 		status VARCHAR(50),
// 		eta VARCHAR(100)
// 	)
// 	`

// 	_, err := DB.Exec(query)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // package main

// // import (
// // 	"database/sql"
// // 	"fmt"
// // 	"os"

// // 	_ "github.com/go-sql-driver/mysql"
// // )

// // func connectDB() (*sql.DB, error) {

// // 	host := os.Getenv("DB_HOST")

// // 	// kalau DB_HOST kosong
// // 	if host == "" {
// // 		host = "localhost"
// // 	}

// // 	user := "root"
// // 	password := "password"
// // 	dbname := "orderdb"

// // 	dsn := fmt.Sprintf(
// // 		"%s:%s@tcp(%s:3306)/%s",
// // 		user,
// // 		password,
// // 		host,
// // 		dbname,
// // 	)

// // 	db, err := sql.Open("mysql", dsn)

// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	err = db.Ping()

// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	return db, nil
// // }



package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	// ambil DB_HOST dari environment variable
	host := os.Getenv("DB_HOST")

	// kalau kosong, pakai localhost
	if host == "" {
		host = "localhost"
	}

	dsn := fmt.Sprintf(
		"root:root@tcp(%s:3306)/tubesdb",
		host,
	)

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

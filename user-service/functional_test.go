//go:build functional

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ===============================
// CONFIG
// ===============================
const (
	dbUser = "root"
	dbPass = "root"
	dbHost = "mysql"
	dbPort = "3306"
	dbName = "tubesdb"
)

func TestUserFlow_Functional(t *testing.T) {

	// ==================================
	// 1. CEK DATABASE BISA DIAKSES
	// ==================================
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatal("database tidak bisa diakses:", err)
	}

	t.Log("DATABASE CONNECTED")

	// ==================================
	// 2. CONNECT APP DATABASE
	// ==================================
	ConnectDB()

	// ==================================
	// 3. START USER SERVICE
	// ==================================
	go func() {

		http.HandleFunc("/register", registerHandler)
		http.HandleFunc("/login", loginHandler)
		http.HandleFunc("/profile", profileHandler)

		http.ListenAndServe(":8081", nil)

	}()

	time.Sleep(2 * time.Second)

	// ==================================
	// 4. REGISTER
	// ==================================
	email := fmt.Sprintf(
		"func%d@mail.com",
		time.Now().UnixNano(),
	)

	respReg, err := http.Post(
		"http://user-service:8081/register",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Name":"Functional",
			"Email":"%s",
			"Password":"123",
			"Role":"customer"
		}`, email))),
	)

	if err != nil {
		t.Fatal(err)
	}

	var reg map[string]interface{}

	json.NewDecoder(respReg.Body).Decode(&reg)

	if respReg.StatusCode != 200 {
		t.Fatalf("register failed: %+v", reg)
	}

	userID := int(reg["user_id"].(float64))

	t.Log("REGISTER SUCCESS")

	// ==================================
	// 5. LOGIN
	// ==================================
	respLogin, err := http.Post(
		"http://user-service:8081/login",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Email":"%s",
			"Password":"123"
		}`, email))),
	)

	if err != nil {
		t.Fatal(err)
	}

	var login map[string]string

	json.NewDecoder(respLogin.Body).Decode(&login)

	token := login["token"]

	if token == "" {
		t.Fatal("login failed")
	}

	t.Log("LOGIN SUCCESS")

	// ==================================
	// 6. GET PROFILE
	// ==================================
	reqProfile, _ := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"http://user-service:8081/profile?id=%d",
			userID,
		),
		nil,
	)

	reqProfile.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	client := &http.Client{}

	respProfile, err := client.Do(reqProfile)

	if err != nil {
		t.Fatal(err)
	}

	var profile map[string]interface{}

	json.NewDecoder(respProfile.Body).Decode(&profile)

	if respProfile.StatusCode != 200 {
		t.Fatalf("get profile failed: %+v", profile)
	}

	t.Log("PROFILE SUCCESS")

	// ==================================
	// 7. CEK USER MASUK DATABASE
	// ==================================
	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE user_id = ?",
		userID,
	).Scan(&count)

	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatal("user tidak masuk database")
	}

	t.Log("USER MASUK DATABASE")
	t.Log("FUNCTIONAL TEST SUCCESS")
}

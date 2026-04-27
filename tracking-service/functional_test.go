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
// Kalau docker compose:
// host = mysql
//
// Kalau lokal:
// host = localhost
//
// Sesuaikan kalau perlu.
const (
	dbUser = "root"
	dbPass = "root"
	dbHost = "mysql"
	dbPort = "3306"
	dbName = "tubesdb" // Gunakan database yang sesuai untuk Tubes
)

func TestInsertTrackingEvent_Functional(t *testing.T) {

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
	// 2. START TRACKING SERVICE LOKAL
	// ==================================
	go func() {
		// Asumsi nama handler kamu adalah insertTrackingEventHandler
		http.HandleFunc("/tracking/event", insertTrackingEventHandler)
		http.ListenAndServe(":8084", nil) // Port 8084 khusus untuk tracking
	}()

	time.Sleep(2 * time.Second)

	// ==================================
	// 3. INSERT TRACKING EVENT
	// ==================================
	// Kita buat resi dinamis untuk test
	resi := fmt.Sprintf("RESI-FUNC-%d", time.Now().UnixNano())

	body := []byte(fmt.Sprintf(`{
		"resi":"%s",
		"lokasi":"Gudang Jakarta",
		"event":"Paket diterima di gudang",
		"timestamp":"2026-04-27T10:00:00Z"
	}`, resi))

	req, _ := http.NewRequest(
		"POST",
		"http://localhost:8084/tracking/event", // Memanggil service lokal yang di-start di atas
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// Test akan FAILED di sini karena insertTrackingEventHandler belum selesai/belum ada
	if resp.StatusCode != 200 {
		t.Fatalf("failed: %+v", result)
	}

	// ==================================
	// 4. CEK DATA MASUK KE DATABASE
	// ==================================
	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM tracking_events WHERE resi = ?",
		resi,
	).Scan(&count)

	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatal("data tracking event tidak masuk database")
	}

	t.Log("TRACKING EVENT MASUK DATABASE")
	t.Log("FUNCTIONAL TEST SUCCESS")
}
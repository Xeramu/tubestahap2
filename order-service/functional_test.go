//go:build functional
// +build functional

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCreateOrder_Functional(t *testing.T) {

	go func() {
		http.HandleFunc("/order", createOrderHandler)
		http.ListenAndServe(":8083", nil)

	}()

	time.Sleep(1 * time.Second)

	// email unik biar tidak bentrok
	email := fmt.Sprintf("func%d@mail.com", time.Now().UnixNano())

	// ====================
	// REGISTER
	// ====================
	respReg, err := http.Post(
		"http://user-service:8081/register",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Name":"Func",
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

	userID := int(reg["user_id"].(float64))

	// ====================
	// LOGIN
	// ====================
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

	// ====================
	// TEST PROFILE DIRECT
	// ====================
	t.Log("Testing profile directly...")

	reqCheck, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("http://user-service:8081/profile?id=%d", userID),
		nil,
	)

	reqCheck.Header.Set("Authorization", "Bearer "+token)

	respCheck, err := http.DefaultClient.Do(reqCheck)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("PROFILE STATUS DIRECT:", respCheck.StatusCode)

	// ====================
	// CREATE ORDER
	// ====================
	body := []byte(fmt.Sprintf(`{
		"user_id":%d,
		"nama_barang":"Laptop",
		"berat":2,
		"dimensi":"10x10",
		"jenis":"Elektronik",
		"alamat_pengirim":"Bandung",
		"alamat_penerima":"Jakarta"
	}`, userID))

	req, _ := http.NewRequest(
		"POST",
		"http://order-service:8083/order",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	t.Log("REGISTER:", reg)
	t.Log("USER ID:", userID)
	t.Log("TOKEN:", token)
	t.Log("ORDER RESPONSE:", result)
	t.Log("STATUS:", resp.StatusCode)

	if resp.StatusCode != 200 {
		t.Fatalf("failed: %+v", result)
	}

	if result["status"] != "created" {
		t.Fatalf("unexpected response: %+v", result)
	}

}

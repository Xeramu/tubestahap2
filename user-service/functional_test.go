//go:build functional

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestUserFlow_Functional(t *testing.T) {

	ConnectDB()

	// ==================================
	// START USER SERVICE
	// ==================================
	go func() {

		http.HandleFunc("/register", registerHandler)
		http.HandleFunc("/login", loginHandler)
		http.HandleFunc("/profile", profileHandler)

		http.ListenAndServe(":8081", nil)

	}()

	time.Sleep(2 * time.Second)

	// ==================================
	// REGISTER
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

	t.Log("REGISTER RESPONSE:", reg)

	if respReg.StatusCode != 200 {
		t.Fatal("register failed")
	}

	userID := int(reg["user_id"].(float64))

	// ==================================
	// LOGIN
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

	t.Log("TOKEN:", token)

	if token == "" {
		t.Fatal("login failed")
	}

	// ==================================
	// GET PROFILE
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

	t.Log("PROFILE:", profile)

	if respProfile.StatusCode != 200 {
		t.Fatal("get profile failed")
	}

	t.Log("FUNCTIONAL TEST SUCCESS")
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var orders []Order
var nextID = 1

// INI BUAT TES LEWAT DOCKER, FUNCTIONAL TES
var UserServiceURL = "http://user-service:8081"

// INI BUAT UNIT TES, LEWAT LOCAL
// var UserServiceURL = "http://localhost:8081"

type Validator interface {
	CheckUser(userID int, token string) bool
}

type RealValidator struct{}

func (v RealValidator) CheckUser(userID int, token string) bool {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("PROFILE ERROR:", err)
		return false
	}

	fmt.Println("PROFILE STATUS:", resp.StatusCode)

	return resp.StatusCode == 200
}

func GenerateResi() string {
	return fmt.Sprintf("LNG-%d", time.Now().Unix())
}

func CalculateETA() string {
	return "2 days"
}

func CreateOrder(req Order, token string, v Validator) (Order, error) {

	if !v.CheckUser(req.UserID, token) {
		return Order{}, errors.New("user invalid")
	}

	if req.Berat <= 0 {
		return Order{}, errors.New("invalid weight")
	}

	req.OrderID = nextID
	req.Resi = GenerateResi()
	req.Status = "created"
	req.ETA = CalculateETA()

	nextID++
	orders = append(orders, req)

	return req, nil
}

func GetOrder(id int) *Order {
	for _, o := range orders {
		if o.OrderID == id {
			return &o
		}
	}
	return nil
}

func UpdateOrderStatus(id int, status string) bool {
	for i := range orders {
		if orders[i].OrderID == id {
			orders[i].Status = status
			return true
		}
	}
	return false
}

func GetETA(id int) string {
	o := GetOrder(id)
	if o == nil {
		return ""
	}
	return o.ETA
}

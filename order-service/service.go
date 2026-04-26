package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	if userID <= 0 {
		return false
	}

	if strings.TrimSpace(token) == "" {
		return false
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
		nil,
	)
	if err != nil {
		return false
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("PROFILE ERROR:", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func GenerateResi() string {
	return fmt.Sprintf("LNG-%d", time.Now().UnixNano())
}

func CalculateETA() string {
	return "2 days"
}

func CreateOrder(req Order, token string, v Validator) (Order, error) {
	if v == nil {
		return Order{}, errors.New("validator required")
	}

	if !v.CheckUser(req.UserID, token) {
		return Order{}, errors.New("user invalid")
	}

	if strings.TrimSpace(req.NamaBarang) == "" {
		return Order{}, errors.New("nama barang required")
	}

	if req.Berat <= 0 {
		return Order{}, errors.New("invalid weight")
	}

	if strings.TrimSpace(req.AlamatPengirim) == "" {
		return Order{}, errors.New("alamat pengirim required")
	}

	if strings.TrimSpace(req.AlamatPenerima) == "" {
		return Order{}, errors.New("alamat penerima required")
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
	for i := range orders {
		if orders[i].OrderID == id {
			return &orders[i]
		}
	}
	return nil
}

func UpdateOrderStatus(id int, status string) bool {
	status = strings.TrimSpace(status)
	if status == "" {
		return false
	}

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

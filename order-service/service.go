// // package main

// // import (
// // 	"errors"
// // 	"fmt"
// // 	"net/http"
// // 	"time"
// // )

// // var orders []Order
// // var nextID = 1

// // // INI BUAT TES LEWAT DOCKER, FUNCTIONAL TES
// // var UserServiceURL = "http://user-service:8081"

// // // INI BUAT UNIT TES, LEWAT LOCAL
// // // var UserServiceURL = "http://localhost:8081"

// // type Validator interface {
// // 	CheckUser(userID int, token string) bool
// // }

// // type RealValidator struct{}

// // func (v RealValidator) CheckUser(userID int, token string) bool {
// // 	req, _ := http.NewRequest(
// // 		"GET",
// // 		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
// // 		nil,
// // 	)

// // 	req.Header.Set("Authorization", "Bearer "+token)

// // 	client := &http.Client{Timeout: 3 * time.Second}
// // 	resp, err := client.Do(req)

// // 	if err != nil {
// // 		fmt.Println("PROFILE ERROR:", err)
// // 		return false
// // 	}

// // 	fmt.Println("PROFILE STATUS:", resp.StatusCode)

// // 	return resp.StatusCode == 200
// // }

// // func GenerateResi() string {
// // 	return fmt.Sprintf("LNG-%d", time.Now().Unix())
// // }

// // func CalculateETA() string {
// // 	return "2 days"
// // }

// // func CreateOrder(req Order, token string, v Validator) (Order, error) {

// // 	if !v.CheckUser(req.UserID, token) {
// // 		return Order{}, errors.New("user invalid")
// // 	}

// // 	if req.Berat <= 0 {
// // 		return Order{}, errors.New("invalid weight")
// // 	}

// // 	req.OrderID = nextID
// // 	req.Resi = GenerateResi()
// // 	req.Status = "created"
// // 	req.ETA = CalculateETA()

// // 	nextID++
// // 	orders = append(orders, req)

// // 	return req, nil
// // }

// // func GetOrder(id int) *Order {
// // 	for _, o := range orders {
// // 		if o.OrderID == id {
// // 			return &o
// // 		}
// // 	}
// // 	return nil
// // }

// // func UpdateOrderStatus(id int, status string) bool {
// // 	for i := range orders {
// // 		if orders[i].OrderID == id {
// // 			orders[i].Status = status
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func GetETA(id int) string {
// // 	o := GetOrder(id)
// // 	if o == nil {
// // 		return ""
// // 	}
// // 	return o.ETA
// // }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// var orders []Order
// var nextID = 1

// // INI BUAT TES LEWAT DOCKER, FUNCTIONAL TES
// var UserServiceURL = "http://user-service:8081"

// // INI BUAT UNIT TES, LEWAT LOCAL
// // var UserServiceURL = "http://localhost:8081"

// type Validator interface {
// 	CheckUser(userID int, token string) bool
// }

// type RealValidator struct{}

// func (v RealValidator) CheckUser(userID int, token string) bool {
// 	req, _ := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
// 		nil,
// 	)

// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{Timeout: 3 * time.Second}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("PROFILE ERROR:", err)
// 		return false
// 	}

// 	fmt.Println("PROFILE STATUS:", resp.StatusCode)

// 	return resp.StatusCode == 200
// }

// func GenerateResi() string {
// 	return ""
// }

// func CalculateETA() string {
// 	return ""
// }

// func CreateOrder(req Order, token string, v Validator) (Order, error) {
// 	return Order{}, nil
// }

// func GetOrder(id int) *Order {
// 	return nil
// }

// func UpdateOrderStatus(id int, status string) bool {
// 	return false
// }

// // func GetETA(id int) string {
// // 	return ""
// // }

package main

import "errors"

var users []User
var nextID = 1

func init() {
	seed := User{
		UserID:   nextID,
		Name:     "seed",
		Email:    "seed@mail.com",
		Password: "hashed",
		Role:     "customer",
	}

	users = append(users, seed)
	nextID++
}

func Register(name, email, password, role string) (User, error) {

	for _, u := range users {
		if u.Email == email {
			return User{}, errors.New("email exists")
		}
	}

	u := User{
		UserID:   nextID,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}

	users = append(users, u)

	// SIMPAN KE DATABASE
	_, err := DB.Exec(
		"INSERT INTO users(name,email,password,role) VALUES(?,?,?,?)",
		u.Name,
		u.Email,
		u.Password,
		u.Role,
	)

	if err != nil {
		return User{}, err
	}

	nextID++

	return u, nil
}

func Login(email, password string) (string, error) {

	for _, u := range users {
		if u.Email == email && u.Password == password {
			return "dummy-token", nil
		}
	}

	return "", errors.New("invalid login")
}

func GetProfile(id int) *User {

	for i := range users {
		if users[i].UserID == id {
			return &users[i]
		}
	}

	return nil
}

func UpdateProfile(id int, alamat string, pref string) bool {

	for i := range users {
		if users[i].UserID == id {
			users[i].Alamat = alamat
			users[i].Preferensi = pref
			return true
		}
	}

	return false
}

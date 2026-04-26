// user-service/service.go
package main

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
	return User{}, nil
}

func Login(email, password string) (string, error) {
	return "", nil
}

func GetProfile(id int) *User {
	return nil
}

func UpdateProfile(id int, alamat string, pref string) bool {
	return false
}

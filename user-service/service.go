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

	hash, _ := HashPassword(password)

	u := User{
		UserID:   nextID,
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     role,
	}

	users = append(users, u)
	nextID++

	return u, nil
}

func Login(email, password string) (string, error) {
	for _, u := range users {
		if u.Email == email && CheckPassword(password, u.Password) {
			return GenerateToken(u.UserID, u.Role), nil
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

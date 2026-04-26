package main

import (
	"errors"
	"strings"
)

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
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)
	role = strings.TrimSpace(strings.ToLower(role))

	if name == "" {
		return User{}, errors.New("name required")
	}

	if email == "" {
		return User{}, errors.New("email required")
	}

	if password == "" {
		return User{}, errors.New("password required")
	}

	if role == "" {
		role = "customer"
	}

	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			return User{}, errors.New("email exists")
		}
	}

	hash, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}

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
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)

	for _, u := range users {
		if strings.EqualFold(u.Email, email) &&
			CheckPassword(password, u.Password) {
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
			users[i].Alamat = strings.TrimSpace(alamat)
			users[i].Preferensi = strings.TrimSpace(pref)
			return true
		}
	}
	return false
}

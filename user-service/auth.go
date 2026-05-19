package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(b), err
}

func CheckPassword(raw, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw)) == nil
}

func GenerateToken(id int, role string) string {
	data := fmt.Sprintf("%d:%s", id, role)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func VerifyToken(t string) bool {
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(t))
	if err != nil {
		return false
	}

	parts := strings.Split(string(decoded), ":")
	return len(parts) == 2
}

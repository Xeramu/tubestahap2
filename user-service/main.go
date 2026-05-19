package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Service Running"))
}

func main() {

	ConnectDB()
	
	http.HandleFunc("/", home)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile", profileHandler)

	fmt.Println("User Service running on :8081")
	http.ListenAndServe(":8081", nil)
}
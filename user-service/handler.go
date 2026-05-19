package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	json.NewDecoder(r.Body).Decode(&req)

	u, err := Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"user_id": u.UserID,
		"name": u.Name,
		"email": u.Email,
		"role": u.Role,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	json.NewDecoder(r.Body).Decode(&req)

	token, err := Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		w.WriteHeader(401)
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if !VerifyToken(token) {
		w.WriteHeader(403)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if r.Method == "GET" {
		u := GetProfile(id)
		if u == nil {
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(u)
		return
	}

	if r.Method == "PUT" {
		var req User
		json.NewDecoder(r.Body).Decode(&req)

		ok := UpdateProfile(id, req.Alamat, req.Preferensi)
		if !ok {
			w.WriteHeader(404)
			return
		}

		w.Write([]byte(`{"message":"updated"}`))
	}
}
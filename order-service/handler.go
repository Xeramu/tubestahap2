package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var validator Validator = RealValidator{}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req Order
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("REQ USERID:", req.UserID)

	token := strings.TrimPrefix(
		r.Header.Get("Authorization"),
		"Bearer ",
	)

	result, err := CreateOrder(req, token, validator)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	order := GetOrder(id)
	if order == nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var req Order
	json.NewDecoder(r.Body).Decode(&req)

	ok := UpdateOrderStatus(id, req.Status)
	if !ok {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}

func etaHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	eta := GetETA(id)
	if eta == "" {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"eta": eta,
	})
}

package main

import (
	"fmt"
	"net/http"
)

func main() {

	ConnectDB()

	http.HandleFunc("/order", createOrderHandler)
	http.HandleFunc("/order/get", getOrderHandler)
	http.HandleFunc("/order/status", updateOrderHandler)
	http.HandleFunc("/order/eta", etaHandler)

	fmt.Println("Order Service running on :8083")
	http.ListenAndServe(":8083", nil)
}

package main

import (
	"log"
	"net/http"
	"order_service/handlers"
)

func main() {
	http.HandleFunc("/order", handlers.GetOrderByUserID)
	log.Println("Order service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

package main

import (
	"log"
	"net/http"
	"user_service/handlers"
)

func main() {
	http.HandleFunc("/user", handlers.GetUserByID)
	log.Println("User service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

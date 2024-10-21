package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order_service/models"
)

func GetOrderByUserID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	for _, order := range models.Orders {
		if fmt.Sprintf("%d", order.UserID) == id {
			user, err := getUserInfo(order.UserID)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			response := map[string]interface{}{
				"order": order,
				"user":  user,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

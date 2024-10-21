package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user_service/models"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	for _, user := range models.Users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

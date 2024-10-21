package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"order_service/models"
)

func getUserInfo(userID int) (*models.User, error) {
	resp, err := http.Get(fmt.Sprintf("http://user_service:8081/user?id=%d", userID))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("User not found")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

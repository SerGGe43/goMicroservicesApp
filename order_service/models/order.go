package models

type Order struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	ItemName string `json:"item_name"`
}

var Orders = []Order{
	{ID: 1, UserID: 1, ItemName: "Laptop"},
	{ID: 2, UserID: 2, ItemName: "Phone"},
}

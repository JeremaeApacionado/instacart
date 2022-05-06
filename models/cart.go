package models

type Cart struct {
	CartId	uint		`json:"cart_id"`
	Products Addtocart	`json:"products"`
	MyUser   User     	`json:"user"`
}
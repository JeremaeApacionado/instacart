package models

type Carts struct {
	Products map[string]interface{} `json:"products"`
	MyUser   Customer                   `json:"user"`
}
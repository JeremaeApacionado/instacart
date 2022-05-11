package models

type User struct {
	UserID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	fullname 	  string `json:"fullname"`
	email   	  string `json:"email"`
	address		  string `json:"address"`
	username	  string `json:"username"`
	password  	  string `json:"password"`
}

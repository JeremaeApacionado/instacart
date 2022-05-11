package models

type User struct {
	UserID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Fullname 	  string `json:"Fullname"`
	Email   	  string `json:"Email"`
	Address		  string `json:"Address"`
	Username	  string `json:"Username"`
	Password  	  string `json:"Password"`
}

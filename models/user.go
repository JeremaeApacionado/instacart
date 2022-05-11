package models

type User struct {
	UserID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Fullname 	  string `json:"fullname"`
	Email   	  string `json:"email"`
	Address		  string `json:"address"`
	Username	  string `json:"username"`
	Password  	  string `json:"password"`
}

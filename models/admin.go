package models

type Admin struct {
	UserID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	NameofStore 	  string `json:"NameofStore"`
	Email   	  string `json:"Email"`
	Address		  string `json:"Address"`
	Username	  string `json:"Username"`
	Password  	  string `json:"Password"`
}

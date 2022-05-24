package models

type Admin struct {
	AdminID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Fullname 	  string `json:"Fullname"`
	Store	 	  string `json:"Name of Store"`
	Email   	  string `json:"Email"`
	Address		  string `json:"Address"`
	Username	  string `json:"Username"`
	Password  	  string `json:"Password"`
}

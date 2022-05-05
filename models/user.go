package models

type User struct {
	ID        	  uint   `json:"id" gorm:"foreignKey"`
	Name 		  string `json:"name"`
	Email   	  string `json:"email"`
	Address		  string `json:"address"`
	Username	  string `json:"username"`
	Password  	  string `json:"password"`
}

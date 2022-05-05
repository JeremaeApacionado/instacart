package models

type User struct {
	ID        uint   `json:"id" gorm:"foreignKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

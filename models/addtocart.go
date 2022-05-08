package models

type Addtocart struct {
	ID              uint `json:"id" gorm:"foreignKey"`
	UserID          uint `json:"Fullname"`
	ProductID       uint `json:"ProductInfo"`
	QuantityOrdered uint `json:"QuantityOrdered"`
}

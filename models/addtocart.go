package models

type Addtocart struct {
	ID  uint   `json:"atc_id" gorm:"foreignKey"`
	UserID      int `json:"Fullname"`
	ProductID	int `json:"Product-Info"`
}


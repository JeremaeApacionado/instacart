package models

type Addtocart struct {
	ID  uint   `json:"atc_id" gorm:"foreignKey"`
	User.ID      int `json:"Fullname"`
	Product.ID	int `json:"Product-Info"`
}


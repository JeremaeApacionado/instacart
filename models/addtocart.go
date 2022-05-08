package models

type Addtocart struct {
	ID  uint   `json:"id" gorm:"foreignKey"`
	UserID      int `json:"UserInfo"`
	ProductID	int `json:"ProductInfo"`
}


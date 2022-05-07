package models

type Addtocart struct {
	ID  uint   `json:"id" gorm:"primaryKey"`
	UserID      int `json:"UserInfo"`
	ProductID	int `json:"ProductInfo"`
}


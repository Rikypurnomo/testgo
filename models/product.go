package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name" validate:"required" gorm:"varchar(255)"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required" gorm:"varchar(255)"`
}

func (Product) TableName() string {
	return "products"
}

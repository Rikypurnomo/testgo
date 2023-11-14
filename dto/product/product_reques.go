package productdto

type CreateProductRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}

type Update1ProductRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

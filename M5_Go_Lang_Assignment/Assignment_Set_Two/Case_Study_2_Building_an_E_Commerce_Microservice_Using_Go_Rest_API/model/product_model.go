package model

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	CategoryID  int     `json:"category_id"`
}

type PaginationQuery struct {
	Page  int `form:"page,default=1" binding:"gte=1"`
	Limit int `form:"limit,default=10" binding:"gte=1,lte=100"`
}
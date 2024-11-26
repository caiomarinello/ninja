package components

type Product struct {
	ProductId int 		`json:"product_id"`
	Name string   		`json:"product_name" validate:"required"`
	Description string 	`json:"product_description" validate:"required"`
	Price float64 		`json:"product_price" validate:"required"`
	Category string 	`json:"product_category" validate:"required"`
	Stock int 			`json:"product_stock" validate:"required"`
}
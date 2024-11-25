package components

type Product struct {
	ProductId int 		`json:"product_id"`
	Name string   		`json:"product_name"`
	Description string 	`json:"product_description"`
	Price float64 		`json:"product_price"`
	Category string 	`json:"product_category"`
	Stock int 			`json:"product_stock"`
}
package repositories

import (
	"database/sql"
	"log"

	comp "ninja/caio/api/components"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Implements the Fetcher interface.
func (p *ProductRepository) FetchById(productId int) (comp.Product, error) {
	var product comp.Product
	query := "SELECT id, product_name, product_description, price, category, stock FROM products WHERE product_id = ?"
	err := p.db.QueryRow(query, productId).Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)
	if err != nil {
		return product, err
	}
	return product, nil
}

// Implements the Fetcher interface.
// Fetches all products from a specific seller.
func (p *ProductRepository) FetchAll() ([]comp.Product, error) {
	var productsSlice []comp.Product
	query := "SELECT id, product_name, product_description, price, category, stock FROM products"
	rows, err := p.db.Query(query)
	if err != nil {
		log.Println("error fetching product rows: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product comp.Product
		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)
		if err != nil {
			log.Println("error scanning product values: ", err)
			return nil, err
		}
		productsSlice = append(productsSlice, product)
	}

	return productsSlice, nil
}
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
	query := "SELECT id, product_name, product_description, price, category, stock FROM products WHERE id = ?"
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

// Implements the Registrar interface:
// it saves a new 'product' in the database.
func (p *ProductRepository) Register(newProduct comp.Product) error {
	query := "INSERT INTO products (product_name, product_description, price, category, stock) VALUES (?, ?, ?, ?, ?)"
	_, err := p.db.Exec(query, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Category, newProduct.Stock)
	if err != nil {
		return err
	}
	return nil
}

// Implements the Updater interface.
func (p *ProductRepository) Update(updatedProduct comp.Product) error {
	query := "UPDATE products SET product_name = ?, product_description = ?, price = ?, category = ?, stock = ? WHERE id = ?"
	_, err := p.db.Exec(query, updatedProduct.Name, updatedProduct.Description, updatedProduct.Price, updatedProduct.Category, updatedProduct.Stock, updatedProduct.ProductId)
	if err != nil {
		return err
	}
	return nil
}

// Implements the Deleter interface.
func (p *ProductRepository) Delete(productId int) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := p.db.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}

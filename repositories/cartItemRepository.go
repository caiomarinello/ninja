package repositories

import (
	"database/sql"
	comp "ninja/caio/api/components"
)

type CartItemRepository struct {
	db *sql.DB
}

// NewCartItemRepository creates a new instance of CartItemRepository.
func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

// Implements the Registrar interface
func (c *CartItemRepository) Register(cartItem comp.CartItem) error {
	query := `INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES (?, ?, ?, ?)`
	_, err := c.db.Exec(query, cartItem.UserId, cartItem.ProductId, cartItem.Quantity, cartItem.TotalPrice)
	if err != nil {
		return err
	}

	return nil
}

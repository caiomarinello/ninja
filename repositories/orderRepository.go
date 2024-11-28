package repositories

import (
	"database/sql"
	"log"
	comp "ninja/caio/api/components"
)

type OrderRepository struct {
	db *sql.DB
}

// NewOrderRepository creates a new instance of OrderRepository.
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Implements the Registrar interface
func (o *OrderRepository) Register(order comp.Order) error {
	query := `INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES (?, ?, ?, ?)`
	_, err := o.db.Exec(query, order.UserId, order.ProductId, order.Quantity, order.TotalPrice)
	if err != nil {
		return err
	}

	return nil
}

// Implements the Fetcher interface.
func (o *OrderRepository) FetchById(orderId int) (comp.Order, error) {
	var order comp.Order
	query := "SELECT id, user_id, product_id, quantity, total_price FROM orders WHERE id = ?"
	err := o.db.QueryRow(query, orderId).Scan(&order.OrderId, &order.UserId, &order.ProductId, &order.Quantity, &order.TotalPrice)
	if err != nil {
		return order, err
	}
	return order, nil
}

// Implements the Fetcher interface.
// Fetches all orders from a specific seller.
func (o *OrderRepository) FetchAll() ([]comp.Order, error) {
	var orderSlice []comp.Order
	query := "SELECT id, user_id, product_id, quantity, total_price FROM orders"
	rows, err := o.db.Query(query)
	if err != nil {
		log.Println("error fetching order rows: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order comp.Order
		err := rows.Scan(&order.OrderId, &order.UserId, &order.ProductId, &order.Quantity, &order.TotalPrice)
		if err != nil {
			log.Println("error scanning order values: ", err)
			return nil, err
		}
		orderSlice = append(orderSlice, order)
	}

	return orderSlice, nil
}
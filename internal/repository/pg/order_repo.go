package pg

import (
	"database/sql"
)

// OrderRepository arayüzü
type OrderRepository interface {
	Create(userID int, productIDs []int) error
	GetByUserID(userID int) ([]Order, error)
}

// Order struct
type Order struct {
	ID     int
	UserID int
	Total  float64
	Status string
}

// OrderRepo struct
type OrderRepo struct {
	db *sql.DB
}

// NewOrderRepo DB ile bağlantı
func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

// Create yeni sipariş ekler
func (r *OrderRepo) Create(userID int, productIDs []int) error {
	// TODO: implement DB insert logic
	return nil
}

// GetByUserID kullanıcıya ait siparişleri döner
func (r *OrderRepo) GetByUserID(userID int) ([]Order, error) {
	// TODO: implement DB select logic
	return []Order{}, nil
}

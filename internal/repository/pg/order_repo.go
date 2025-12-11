package pg

import (
	"database/sql"
)

type OrderRepository interface {
	Create(userID int, productIDs []int) error
	GetByUserID(userID int) ([]Order, error)
}

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepository {
	return &orderRepo{db: db}
}

// Order modeli
type Order struct {
	ID        int
	UserID    int
	Total     float64
	Status    string
	CreatedAt string
	UpdatedAt string
}

// Create yeni sipariş ekler
func (r *orderRepo) Create(userID int, productIDs []int) error {
	// TODO: implement
	return nil
}

// GetByUserID kullanıcıya ait siparişleri döner
func (r *orderRepo) GetByUserID(userID int) ([]Order, error) {
	// TODO: implement
	return nil, nil
}

package pg

import (
	"database/sql"

	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
)

type cartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) CartRepository {
	return &cartRepo{db: db}
}

func (r *cartRepo) AddItem(item *models.CartItem) error {
	query := `INSERT INTO cart (user_id, product_id, quantity, price) VALUES ($1,$2,$3,$4) RETURNING id`
	return r.db.QueryRow(query, item.UserID, item.ProductID, item.Quantity, item.Price).Scan(&item.ID)
}

func (r *cartRepo) UpdateItem(item *models.CartItem) error {
	query := `UPDATE cart SET quantity=$1, price=$2 WHERE id=$3`
	_, err := r.db.Exec(query, item.Quantity, item.Price, item.ID)
	return err
}

func (r *cartRepo) DeleteItem(id int) error {
	_, err := r.db.Exec(`DELETE FROM cart WHERE id=$1`, id)
	return err
}

func (r *cartRepo) GetCartByUser(userID int) ([]*models.CartItem, error) {
	rows, err := r.db.Query(`SELECT id,user_id,product_id,quantity,price FROM cart WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*models.CartItem{}
	for rows.Next() {
		item := &models.CartItem{}
		if err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *cartRepo) ClearCart(userID int) error {
	_, err := r.db.Exec(`DELETE FROM cart WHERE user_id=$1`, userID)
	return err
}

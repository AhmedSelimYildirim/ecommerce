package pg

import (
	"database/sql"
	"time"

	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, stock, created_at, updated_at)
			  VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.Stock, time.Now(), time.Now()).Scan(&product.ID)
	return err
}

func (r *productRepo) GetByID(id int) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT id,name,description,price,stock,created_at,updated_at FROM products WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepo) Update(product *models.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4, updated_at=$5 WHERE id=$6`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, time.Now(), product.ID)
	return err
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *productRepo) GetAll() ([]*models.Product, error) {
	rows, err := r.db.Query(`SELECT id,name,description,price,stock,created_at,updated_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

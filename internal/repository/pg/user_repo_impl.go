package pg

import (
	"database/sql"
	"errors"

	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	query := `INSERT INTO users (name, email, password, role, created_at)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role, user.CreatedAt).Scan(&user.ID)
	return err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE email=$1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

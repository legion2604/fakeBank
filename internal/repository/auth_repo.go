package repository

import (
	"database/sql"
	"fakeBank/internal/models"
	"fakeBank/pkg/errors"
	"time"
)

type AuthRepository interface {
	Login(email string) (models.GetUser, string, error)
	Signup(firstName, lastName, email, password, phone string) (bool, int, time.Time, error)
	Me(userId int) (models.GetUser, error)
}
type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

// methods

func (r *authRepository) Login(email string) (models.GetUser, string, error) {
	var res models.GetUser
	var hashedPassword string
	err := r.db.QueryRow("SELECT id, first_name, last_name, email,password, phone, created_at FROM users WHERE email=$1", email).Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &hashedPassword, &res.Phone, &res.CreatedAt)
	if err != nil {
		return models.GetUser{}, "", errors.ErrEmailExists
	}
	return res, hashedPassword, nil
}

func (r *authRepository) Signup(firstName, lastName, email, password, phone string) (bool, int, time.Time, error) {
	var createdAt time.Time
	var userId int
	err := r.db.QueryRow("INSERT INTO users (first_name, last_name, email, password, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id,created_at", firstName, lastName, email, password, phone).Scan(&userId, &createdAt)
	if err != nil {
		return false, 0, time.Time{}, nil
	}
	return true, userId, createdAt, nil
}

func (r *authRepository) Me(userId int) (models.GetUser, error) {
	var user models.GetUser
	err := r.db.QueryRow("SELECT first_name, last_name, email, phone, created_at FROM users WHERE id=$1", userId).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.CreatedAt)
	user.Id = userId
	if err != nil {
		return models.GetUser{}, errors.ErrUserNotFound
	}
	return user, nil
}

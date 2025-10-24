package repository

import (
	"database/sql"
	"fakeBank/internal/models"
)

type UserRepository interface {
	GetProfile(userId int) (models.GetUser, error)
	GetUserPassword(userId int) (string, error)
	ChangePassword(userId int, newPassword string) (bool, error)
	DeleteUser(userId int) (bool, error)
	UpdateUser(userId int, user models.UpdateUserJson) (models.GetUser, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (u userRepo) GetUserPassword(userId int) (string, error) {
	var userPass string
	err := u.db.QueryRow("SELECT password FROM users WHERE id=$1", userId).Scan(&userPass)
	if err != nil {
		return "", err
	}
	return userPass, nil
}

func (u userRepo) ChangePassword(userId int, newPassword string) (bool, error) {
	_, err := u.db.Exec("UPDATE users SET password=$1 WHERE id=$2", newPassword, userId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u userRepo) GetProfile(userId int) (models.GetUser, error) {
	var user models.GetUser
	err := u.db.QueryRow("SELECT first_name, last_name, email, phone, created_at FROM users WHERE id = $1", userId).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.CreatedAt)
	user.Id = userId
	if err != nil {
		return models.GetUser{}, err
	}
	return user, nil
}

func (u userRepo) DeleteUser(userId int) (bool, error) {
	_, err := u.db.Exec("DELETE FROM users WHERE id=$1", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u userRepo) UpdateUser(userId int, user models.UpdateUserJson) (models.GetUser, error) {
	var updatedUser models.GetUser
	err := u.db.QueryRow("UPDATE users SET first_name=$1, last_name=$2, phone=$3 WHERE id=$4 RETURNING id, first_name, last_name, email, phone, created_at", user.FirstName, user.LastName, user.Phone, userId).Scan(&updatedUser.Id, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Email, &updatedUser.Phone, &updatedUser.CreatedAt)
	if err != nil {
		return models.GetUser{}, err
	}
	return updatedUser, nil
}

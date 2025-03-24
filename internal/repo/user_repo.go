package repo

import (
	"context"
	"database/sql"
	"fmt"
	"go-brain/internal/model"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO users (name, password) VALUES (?, ?)",
		user.Name, user.Password)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.DB.QueryRowContext(ctx,
		"SELECT id, name, password FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	_, err := r.DB.ExecContext(ctx,
		"UPDATE users SET name = ?, password = ? WHERE id = ?",
		user.Name, user.Password, user.ID)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx,
		"DELETE FROM users WHERE id = ?", id)
	return err
}

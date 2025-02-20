package repo

import (
	"context"


	"go-brain/internal/model"

	"github.com/uptrace/bun"
)

type UserRepo struct {
	DB *bun.DB
}

func NewUserRepo(db *bun.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	_, err := r.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.DB.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	return err
}

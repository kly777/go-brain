package internal

import (
	"context"

	"github.com/uptrace/bun"
)

type UserRepo struct {
	DB *bun.DB
}

func NewUserRepo(db *bun.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *User) error {
	_, err := r.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*User, error) {
	user := new(User)
	err := r.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *User) error {
	_, err := r.DB.NewUpdate().Model(user).WherePK().Exec(ctx)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.NewDelete().Model(&User{}).Where("id = ?", id).Exec(ctx)
	return err
}

type ThingRepo struct {
	DB *bun.DB
}

func NewThingRepo(db *bun.DB) *ThingRepo {
	return &ThingRepo{DB: db}
}

func (r *ThingRepo) Create(ctx context.Context, thing *Thing) error {
	_, err := r.DB.NewInsert().Model(thing).Exec(ctx)
	return err
}

func (r *ThingRepo) GetByID(ctx context.Context, id int64) (*Thing, error) {
	thing := new(Thing)
	err := r.DB.NewSelect().Model(thing).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return thing, nil
}

func (r *ThingRepo) Update(ctx context.Context, thing *Thing) error {
	_, err := r.DB.NewUpdate().Model(thing).WherePK().Exec(ctx)
	return err
}

func (r *ThingRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.NewDelete().Model(&Thing{}).Where("id = ?", id).Exec(ctx)
	return err
}
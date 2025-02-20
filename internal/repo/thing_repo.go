package repo

import (
	"context"
	"github.com/uptrace/bun"
	"go-brain/internal/model"
)

type ThingRepo struct {
	DB *bun.DB
}

func NewThingRepo(db *bun.DB) *ThingRepo {
	return &ThingRepo{DB: db}
}

func (r *ThingRepo) Create(ctx context.Context, thing *model.Thing) error {
	_, err := r.DB.NewInsert().Model(thing).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ThingRepo) GetByID(ctx context.Context, id int64) (*model.Thing, error) {
	var thing model.Thing
	err := r.DB.QueryRowContext(ctx, "SELECT id, name FROM things WHERE id = ?", id).Scan(&thing.ID, &thing.Name)
	if err != nil {
		return nil, err
	}
	return &thing, nil
}

func (r *ThingRepo) Update(ctx context.Context, thing *model.Thing) error {
	_, err := r.DB.NewUpdate().Model(thing).WherePK().Exec(ctx)
	return err
}

func (r *ThingRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.NewDelete().Model(&model.Thing{}).Where("id = ?", id).Exec(ctx)
	return err
}

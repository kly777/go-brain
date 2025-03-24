package repo

import (
	"context"
	"database/sql"
	"go-brain/internal/model"
)

type ThingRepo struct {
	DB *sql.DB
}

func NewThingRepo(db *sql.DB) *ThingRepo {
	return &ThingRepo{DB: db}
}

func (r *ThingRepo) Create(ctx context.Context, thing *model.Thing) error {
	stmt, err := r.DB.Prepare("INSERT INTO things (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, thing.Name)
	return err
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
	stmt, err := r.DB.Prepare("UPDATE things SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, thing.Name, thing.ID)
	return err
}

func (r *ThingRepo) Delete(ctx context.Context, id int64) error {
	stmt, err := r.DB.Prepare("DELETE FROM things WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}

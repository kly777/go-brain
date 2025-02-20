package service

import (
	"context"
	"errors"
	"go-brain/internal/model"
	"go-brain/internal/repo"
)


type ThingService struct {
	ThingRepo *repo.ThingRepo
}

func NewThingService(thingRepo *repo.ThingRepo) *ThingService {
	return &ThingService{ThingRepo: thingRepo}
}

func (s *ThingService) Create(ctx context.Context, thing *model.Thing) error {
	return s.ThingRepo.Create(ctx, thing)
}

func (s *ThingService) GetByID(ctx context.Context, id int64) (*model.Thing, error) {
	return s.ThingRepo.GetByID(ctx, id)
}

func (s *ThingService) Update(ctx context.Context, thing *model.Thing) error {
	if thing.ID == 0 {
		return errors.New("thing ID is required")
	}
	return s.ThingRepo.Update(ctx, thing)
}

func (s *ThingService) Delete(ctx context.Context, id int64) error {
	return s.ThingRepo.Delete(ctx, id)
}

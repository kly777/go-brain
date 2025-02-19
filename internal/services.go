package internal

import (
	"context"
	"errors"
)

type UserService struct {
	UserRepo *UserRepo
}

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, user *User) error {
	return s.UserRepo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.UserRepo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *User) error {
	if user.ID == 0 {
		return errors.New("user ID is required")
	}
	return s.UserRepo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.UserRepo.Delete(ctx, id)
}

type ThingService struct {
	ThingRepo *ThingRepo
}

func NewThingService(thingRepo *ThingRepo) *ThingService {
	return &ThingService{ThingRepo: thingRepo}
}

func (s *ThingService) Create(ctx context.Context, thing *Thing) error {
	return s.ThingRepo.Create(ctx, thing)
}

func (s *ThingService) GetByID(ctx context.Context, id int64) (*Thing, error) {
	return s.ThingRepo.GetByID(ctx, id)
}

func (s *ThingService) Update(ctx context.Context, thing *Thing) error {
	if thing.ID == 0 {
		return errors.New("thing ID is required")
	}
	return s.ThingRepo.Update(ctx, thing)
}

func (s *ThingService) Delete(ctx context.Context, id int64) error {
	return s.ThingRepo.Delete(ctx, id)
}
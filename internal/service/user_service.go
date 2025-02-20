package service

import (
	"context"
	"errors"
	"go-brain/internal/model"
	"go-brain/internal/repo"
)


type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	return s.UserRepo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return s.UserRepo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *model.User) error {
	if user.ID == 0 {
		return errors.New("user ID is required")
	}
	return s.UserRepo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.UserRepo.Delete(ctx, id)
}
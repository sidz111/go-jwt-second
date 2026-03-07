package service

import (
	"context"
	"errors"

	"github.com/sidz111/jwt-second-project/models"
	"github.com/sidz111/jwt-second-project/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (uint, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) (uint, error) {
	if err := ValidateUser(user); err != nil {
		return 0, err
	}
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(encryptPassword)
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAllUsers(ctx)
}
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid user id")
	}
	return s.repo.DeleteUser(ctx, id)
}
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	if user.ID <= 0 {
		return errors.New("invalid user id")
	}
	if err := ValidateUser(user); err != nil {
		return err
	}
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(encryptPassword)
	return s.repo.UpdateUser(ctx, user)
}

func ValidateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

package repository

import (
	"context"
	"errors"

	"github.com/sidz111/jwt-second-project/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (int, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (int, error) {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}
func (r *userRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	// result := r.db.Where("id=?", id).First(&models.User{}, id)
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	result := r.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user)
	// result := r.db.WithContext(ctx).Save(&user) ==> we can use above line instead of the current line, both will work fine
	if result.Error != nil {
		return result.Error
	}
	return nil
}

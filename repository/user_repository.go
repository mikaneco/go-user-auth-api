package repository

import (
	"goapi/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) (*model.User, error)
	IsExistByEmail(email string) bool
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) FindById(id string) (*model.User, error) {
	var user model.User
	result := repository.db.First(&user, id)
	return &user, result.Error
}

func (repository *userRepository) IsExistByEmail(email string) bool {
	var user model.User
	result := repository.db.Where("email = ?", email).First(&user)
	return result.RowsAffected > 0
}

func (repository *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := repository.db.First(&user, "email = ?", email)
	return &user, result.Error
}

func (repository *userRepository) Create(user *model.User) error {
	result := repository.db.Create(user)
	return result.Error
}

func (repository *userRepository) Update(user *model.User) error {
	result := repository.db.Save(user)
	return result.Error
}

func (repository *userRepository) Delete(email string) error {
	result := repository.db.Delete(&model.User{}, "email = ?", email)
	return result.Error
}

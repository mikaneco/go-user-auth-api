package service

import (
	"fmt"
	"goapi/model"
	"goapi/repository"
)

type UserService interface {
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(email string) error
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{UserRepository: userRepository}
}

func (service *userService) FindById(id string) (*model.User, error) {
	return service.UserRepository.FindById(id)
}

func (service *userService) FindByEmail(email string) (*model.User, error) {
	return service.UserRepository.FindByEmail(email)
}

func (service *userService) Create(user *model.User) (*model.User, error) {
	// 同じメールアドレスのユーザーが存在するかチェック
	if service.UserRepository.IsExistByEmail(user.Email) {
		return nil, fmt.Errorf("email %s is already registered", user.Email)
	}

	// パスワードをハッシュ化
	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	// ユーザーを作成
	if err := service.UserRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) Update(user *model.User) (*model.User, error) {
	return user, service.UserRepository.Update(user)
}

func (service *userService) Delete(email string) error {
	return service.UserRepository.Delete(email)
}

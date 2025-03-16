package service

import "go-ecommerce-app/internal/domain"

type UserService struct {
}

func (s UserService) Signup(input any) (string, error) {
	return "", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) GetVerificationCode(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) DeleteProfile(id uint) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(id uint) ([]domain.User, error) {
	return nil, nil
}

func (s UserService) GetOrderByID(id uint, uId uint) (interface{}, error) {
	return nil, nil
}

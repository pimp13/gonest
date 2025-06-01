package users

import (
	"errors"
	"github.com/pimp13/gonest/config"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	// dependence injection
	config *config.Config
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		config: cfg,
	}
}

func (us *UserService) FindAll() ([]User, error) {
	return []User{
		{ID: 1, Name: "Pouya", Email: "pouya@email.com"},
		{ID: 2, Name: "Ali", Email: "ali@email.com"},
		{ID: 3, Name: "Sara", Email: "sara@email.com"},
		{ID: 4, Name: "Ariya", Email: "ariya@email.com"},
	}, nil
}

func (s *UserService) FindByID(id uint) (*User, error) {
	if id == 1 {
		return &User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil
	}
	return nil, errors.New("user not found")
}

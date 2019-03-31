package service

import (
	"github.com/jinzhu/gorm"
	"go-webapi-for-gae-study/backend/model"
)

// User ...
type User interface {
	CreateUser(m *model.User) (string, error)
	GetUser(id string) (*model.User, error)
	UpdateUser(m *model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userService struct {
	db *gorm.DB
}

// NewUser ...
func NewUser(db *gorm.DB) User {
	return &userService{db: db}
}

// CreateUser ...
func (n *userService) CreateUser(m *model.User) (string, error) {
	return model.NewUserDao(n.db).CreateUser(m)
}

// GetUser ...
func (n *userService) GetUser(id string) (*model.User, error) {
	return model.NewUserDao(n.db).GetUser(id)
}

// UpdateUser ...
func (n *userService) UpdateUser(m *model.User) (*model.User, error) {
	return model.NewUserDao(n.db).UpdateUser(m)
}

// DeleteUser ...
func (n *userService) DeleteUser(id string) error {
	return model.NewUserDao(n.db).DeleteUser(id)
}

package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Login ...
type Login struct {
	ID       string `gorm:"column:id;primary_key"`
	Email    string `gorm:"column:email;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

// IsDto ...
func (n *Login) IsDto() bool { return true }

// Logout ...
type Logout struct {
	ID       string `gorm:"column:id;primary_key"`
	Email    string `gorm:"column:email;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

// IsDto ...
func (n *Logout) IsDto() bool { return true }

// AuthDao ...
type AuthDao interface {
	Login(l *Login) (string, error)
	Logout(l *Logout) error
}

type authDao struct {
	db *gorm.DB
}

// NewAuthDao ...
func NewAuthDao(db *gorm.DB) AuthDao {
	return &authDao{
		db: db,
	}
}

// Login ...
func (n *authDao) Login(l *Login) (string, error) {
	jwtToken := strings.Replace(uuid.New().String(), "-", "", -1)
	return jwtToken, nil
}

// Logout ...
func (n *authDao) Logout(l *Logout) error {
	return nil
}

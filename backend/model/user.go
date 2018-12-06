package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strings"
)

// User ...
type User struct {
	ID string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

// IsDto ...
func (n *User) IsDto() bool { return true }

// UserDao ...
type UserDao interface {
	CreateUser(m *User) (string, error)
	GetUser(id string) (*User, error)
	UpdateUser(m *User) (*User, error)
	DeleteUser(id string) error
}

type userDao struct {
	db *gorm.DB
}

// NewUserDao ...
func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

// CreateUser ...
func (n *userDao) CreateUser(m *User) (string, error) {
	m.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	n.db.Create(&m)
	if n.db.Error != nil {
		return "", n.db.Error
	}
	return m.ID, nil
}

// GetUser ...
func (n *userDao) GetUser(id string) (*User, error) {
	result := &User{}
	n.db.Where(&User{ID: id}).First(result)
	if n.db.Error != nil {
		return nil, n.db.Error
	}
	return result, nil
}

// UpdateUser ...
func (n *userDao) UpdateUser(m *User) (*User, error) {
	nowRecord, err := n.GetUser(m.ID)
	if err != nil {
		return nil, err
	}
	if nowRecord.ID == "" {
		return nowRecord, nil
	}
	result := &User{
		ID:        nowRecord.ID,
		Name:      m.Name,
	}
	n.db.Save(result)
	if n.db.Error != nil {
		return nil, n.db.Error
	}
	return result, nil
}

// DeleteUser ...
func (n *userDao) DeleteUser(id string) error {
	n.db.Delete(&User{ID: id})
	if n.db.Error != nil {
		return n.db.Error
	}
	return nil
}

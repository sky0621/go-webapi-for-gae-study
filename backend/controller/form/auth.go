package form

import (
	"go-webapi-for-gae-study/backend/model"
)

// Login ...
type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ParseToDto ...
func (n *Login) ParseToDto() *model.Login {
	// フォーマット変換等は、ここで吸収
	return &model.Login{
		Email:    n.Email,
		Password: n.Password,
	}
}

// Logout ...
type Logout struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ParseToDto ...
func (n *Logout) ParseToDto() *model.Logout {
	// フォーマット変換等は、ここで吸収
	return &model.Logout{
		Email:    n.Email,
		Password: n.Password,
	}
}

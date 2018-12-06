package form

import (
	"time"
)

// REF: https://github.com/go-playground/validator/blob/v9/_examples/simple/main.go

// CreateUser ...
type CreateUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Age  		int    `json:"age" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ParseToDto ...
func (n *CreateUser) ParseToDto() *model.User {
	// フォーマット変換等は、ここで吸収
	return &model.User{
		ID:        n.ID,
		Code:      n.Code,
		Sentence:  n.Sentence,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

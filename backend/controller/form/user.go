package form

import (
	"go-webapi-for-gae-study/backend/model"
)

// User ...
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ParseToDto ...
func (n *User) ParseToDto() *model.User {
	// フォーマット変換等は、ここで吸収
	return &model.User{
		ID:   n.ID,
		Name: n.Name,
	}
}

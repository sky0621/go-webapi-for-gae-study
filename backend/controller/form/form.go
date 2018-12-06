package form

import "go-webapi-for-gae-study/backend/model"

// Form ...
type Form interface {
	ParseToDto() model.Dto
}

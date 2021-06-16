package userin

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type Search struct {
	Type string `json:"type" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Type:  input.Type,
		Value: input.Value,
	}
}

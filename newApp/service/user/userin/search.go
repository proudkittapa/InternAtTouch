package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type Search struct {
	Type string `json:"type" validate:"required"`
	// ID        string `json:"id" validate:"required"`
	Value string `json:"value" validate:"required"`
	// Tel       string `json:"tel" validate:"required"`
} // @Name StaffCreateInput

func MakeTestSearchInput() (input *Search) {
	return &Search{
		Type: "test",
		// ID:        "test",
		Value: "test",
		// Tel:       "test",
	}
}

func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Type:  input.Type,
		Value: input.Value,
	}
}

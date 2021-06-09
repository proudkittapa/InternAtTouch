package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type Search struct {
	Type string `json:"ID" validate:"required"`
	// ID        string `json:"id" validate:"required"`
	Value string `json:"name" validate:"required"`
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
		// CompanyID: input.CompanyID,
		// ID:        input.ID,
		// Name: input.Name,

		// Tel:       input.Tel,
		// CreatedAt: carbon.Now().Unix(),
		// UpdatedAt: carbon.Now().Unix(),
	}
}

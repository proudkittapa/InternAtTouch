package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"touch/domain"
)

type CreateInput struct {
	ID string `json:"ID" validate:"required"`
	// ID        string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	// Tel       string `json:"tel" validate:"required"`
} // @Name StaffCreateInput

func MakeTestCreateInput() (input *CreateInput) {
	return &CreateInput{
		ID: "test",
		// ID:        "test",
		Name: "test",
		// Tel:       "test",
	}
}

func CreateInputToUserDomain(input *CreateInput) (user *domain.User) {
	return &domain.User{
		// CompanyID: input.CompanyID,
		// ID:        input.ID,
		Name: input.Name,
		// Tel:       input.Tel,
		// CreatedAt: carbon.Now().Unix(),
		// UpdatedAt: carbon.Now().Unix(),
	}
}

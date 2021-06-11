package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

// type CreateInput struct {
// 	ID string `json:"ID" validate:"required"`
// 	// ID        string `json:"id" validate:"required"`
// 	Name string `json:"name" validate:"required"`
// 	// Tel       string `json:"tel" validate:"required"`
// } // @Name StaffCreateInput

type CreateInput struct {
	ID         string   `json:"id"`
	Name       string   `son:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
}

func MakeTestCreateInput() (input *CreateInput) {
	return &CreateInput{
		ID:   "test",
		Name: "test",
	}
}

func CreateInputToUserDomain(input *CreateInput) (user *domain.InsertQ) {
	return &domain.InsertQ{
		ID:         input.ID,
		Name:       input.Name,
		ActualName: input.ActualName,
		Gender:     input.Gender,
		BirthDate:  input.BirthDate,
		Height:     input.Height,
		SuperPower: input.SuperPower,
		Alive:      input.Alive,
	}
}

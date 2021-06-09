package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type UpdateInput struct {
	ID         string   `json:"id"`
	Name       string   `son:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
} // @Name StaffCreateInput

func MakeTestUpdateInput() (input *UpdateInput) {
	return &UpdateInput{
		ID: "test",
		// ID:        "test",
		Name: "test",
		// Tel:       "test",
	}
}

func CreateInputToUpdateDomain(input *CreateInput) (user *domain.UpdateQ) {
	return &domain.UpdateQ{
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

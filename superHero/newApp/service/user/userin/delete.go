package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteQ) {
	return &domain.DeleteQ{
		ID: input.ID,
	}
}

func (input *CreateInput)DeleteInputToUserDomain() (user *domain.DeleteQ) {
	return &domain.DeleteQ{
		ID:             input.ID,
	}
}
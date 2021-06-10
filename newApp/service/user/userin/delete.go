package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
} // @Name StaffCreateInput

func MakeTestDeleteInput() (input *DeleteInput) {
	return &DeleteInput{
		// Type: "test",
		ID: "test",
		// Value: "test",
		// Tel:       "test",
	}
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteQ) {
	return &domain.DeleteQ{
		ID: input.ID,
	}
}

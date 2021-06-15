package userin

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type ViewInput struct {
	ID string `json:"id"`
}

func ViewInputToUserDomain(input *ViewInput) (user *domain.ViewQ) {
	return &domain.ViewQ{
		ID: input.ID,
	}
}

package userin

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type ViewAllInput struct {
	PerPage int
	Page    int
}

func ViewAllInputToUserDomain(input *ViewAllInput) (user *domain.ViewByPageQ) {
	return &domain.ViewByPageQ{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}

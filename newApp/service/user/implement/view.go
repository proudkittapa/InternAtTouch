package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) View(ctx context.Context, input *userin.ViewInput) (a domain.InsertQ, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return a, err
	}

	user := userin.ViewInputToUserDomain(input)

	a, err = impl.repo.View(ctx, user.ID)
	if err != nil {
		return a, err
	}

	return a, nil
}

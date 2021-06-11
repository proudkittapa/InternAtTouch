package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) ViewAll(ctx context.Context, input *userin.ViewAllInput) (a []domain.InsertQ, err error) {
	// err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", err
	// }

	user := userin.ViewAllInputToUserDomain(input)

	a, err = impl.repo.ViewAll(ctx, user.PerPage, user.Page)
	if err != nil {
		return a, err
	}

	return a, nil
}

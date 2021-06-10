package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	// err = impl.validator.Validate(input)
	if err != nil {
		return "", err
	}

	user := userin.UpdateInputToUserDomain(input)

	err = impl.repo.Update(ctx, user, user.ID)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Search(ctx context.Context, input *userin.Search) (result string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", err
	}

	user := userin.SearchInputToUserDomain(input)

	_, err = impl.repo.Search(ctx, user)
	if err != nil {
		return "", err
	}

	return user.Value, nil
}

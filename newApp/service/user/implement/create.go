package implement

import (
	"context"
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	// err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", err
	// }

	user := userin.CreateInputToUserDomain(input)
	fmt.Println("user input create:", user)

	err = impl.repo.Create(ctx, user)
	// fmt.Println("output create:", user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

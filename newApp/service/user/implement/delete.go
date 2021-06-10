package implement

import (
	"context"
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	// err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", err
	// }

	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	err = impl.repo.Delete(ctx, user)
	fmt.Println("output del:", user)
	fmt.Println("err del:", err)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

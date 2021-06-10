package implement

import (
	"context"
	"fmt"

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
	fmt.Println("user input view: ", user)
	fmt.Println("user.ID: ", user.ID)
	a, err = impl.repo.View(ctx, user.ID)
	fmt.Println("out put view: ", a)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}

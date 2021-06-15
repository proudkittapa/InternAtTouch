package implement

import (
	"context"
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	user := userin.UpdateInputToUserDomain(input)

	err = impl.repo.Update(ctx, user, user.ID)
	if err != nil {
		// fmt.Println("er")
		return "", err
	}

	return user.Name, nil
}

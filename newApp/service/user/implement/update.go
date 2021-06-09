package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", util.ValidationCreateErr(err)
	// }

	user := userin.CreateInputToUserDomain(input)

	// _, err = impl.repo.Create(ctx, company)
	// if err != nil {
	// 	return "", util.RepoCreateErr(err)
	// }

	return user.Name, nil
}

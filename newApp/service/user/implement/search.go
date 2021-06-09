package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Search(ctx context.Context, input *userin.Search) (ID string, err error) {
	user := userin.SearchInputToUserDomain(input)

	_, err = impl.repo.Create(ctx, user)
	// if err != nil {
	// 	return "", util.RepoCreateErr(err)
	// }
	//	switch field{
	//	case "name", "actual_name", "gender", "super_power":
	//SearchByField()
	//	case "both_name":
	//SearchByBothName
	//	default:
	//		return
	//	}
	return user.Name, nil
}

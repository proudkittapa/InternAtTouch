package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Search(ctx context.Context, input *userin.CreateInput) (ID string, err error) {

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

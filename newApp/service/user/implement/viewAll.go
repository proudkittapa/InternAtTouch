package implement

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (impl *implementation) ViewAll(ctx context.Context, input *userin.ViewAllInput) (a []domain.InsertQ, err error) {
	user := userin.ViewAllInputToUserDomain(input)

	a, err = impl.repo.ViewAll(ctx, user.PerPage, user.Page)
	if err != nil {
		return a, err
	}

	return a, nil
}

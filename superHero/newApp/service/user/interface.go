package user

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

type Service interface {
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Search(ctx context.Context, input *userin.Search) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	View(ctx context.Context, input *userin.ViewInput) (a domain.InsertQ, err error)
	ViewAll(ctx context.Context, input *userin.ViewAllInput) (a []domain.InsertQ, err error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
}

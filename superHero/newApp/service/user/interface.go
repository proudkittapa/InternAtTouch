package user

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

//go:generate mockery --name=Service
type Service interface {
	// List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.StaffView, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Search(ctx context.Context, input *userin.Search) (map[string]interface{}, error)
	// Read(ctx context.Context, input *staffin.ReadInput) (staff *out.StaffView, err error)
	// Update(ctx context.Context, input *staffin.UpdateInput) (err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	View(ctx context.Context, input *userin.ViewInput) (map[string]interface{}, error)
	ViewAll(ctx context.Context, input *userin.ViewAllInput) (map[string]interface{}, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
}

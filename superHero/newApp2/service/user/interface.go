package user

import (
	"context"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

type Service interface {
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
}

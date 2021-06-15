package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	err = impl.repo.Delete(ctx, user.ID)
	fmt.Println("output del:", user)
	fmt.Println("err del:", err)

	if err != nil {
		return "", err
	}

	return user.ID, err
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
	return impl.MsgSender("responseDelete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDeleteResponse,
		ID:             input.ID,
	})
}

package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (impl *implementation) MsgReceiver(ctx context.Context, msg []byte) (err error) {
	msgInput := &userin.MsgBrokerCreate{}
	err = json.Unmarshal(msg, msgInput)
	if err != nil {
		return err
	}

	//msgAuthInput := &
	//err = json.Unmarshal(msg, msgInput)

	fmt.Println(string(msg))
	switch msgInput.Action {
	case msgbrokerin.ActionCreate:
		err = impl.receiveCreateAction(ctx, msgInput)
		if err != nil {
			return err
		}

	//case msgbrokerin.ActionUpdate:
	//	//    TODO Update Users
	//	fmt.Println(fmt.Sprintf("%s has updated", msgInput.FirstName.En))
	//case msgbrokerin.ActionDelete:
	//	//    TODO Delete Users
	//	fmt.Println(fmt.Sprintf("%s has deleted", msgInput.FirstName.En))

	}

	return
}

func (impl *implementation) receiveCreateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToCreateInput()
	domainUser := input.CreateInputToUserDomain()
	err = impl.repo.Create(ctx, domainUser)
	if err != nil {
		return err
	}

	return nil
}

package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

func (impl implementation) newHandler(topic msgbrokerin.TopicMsgBroker) (handler common.Handler) {
	fmt.Println("enter newHandler", topic)
	return func(ctx context.Context, msg []byte) {
		fmt.Println("msg", msg)
		var err error
		switch topic {
		case msgbrokerin.TopicCreate:
			fmt.Println("topic create is received")
			//var input *userin.CreateInput
			//err = json.Unmarshal(msg, input)
			//impl.usrService.Create(context.Background(), input)
			err = impl.usrService.MsgReceiver(ctx, msg)
		default:
			fmt.Println(string(msg), "with default")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
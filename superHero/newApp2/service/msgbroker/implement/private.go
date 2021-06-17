package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

func (impl implementation) newHandler(topic msgbrokerin.TopicMsgBroker) (handler common.Handler) {
	fmt.Println("enter newHandler", topic)
	return func(ctx context.Context, msg []byte) {
		//fmt.Println("msg", msg)
		var err error
		switch topic {
		case msgbrokerin.TopicCreate:
			fmt.Println("topic create is received")
			err = impl.usrService.MsgReceiver(ctx, msg)
		case msgbrokerin.TopicUpdate:
			fmt.Println("topic update is received")
			err = impl.usrService.MsgReceiver(ctx, msg)
		case msgbrokerin.TopicDelete:
			fmt.Println("topic delete is received")
			err = impl.usrService.MsgReceiver(ctx, msg)
		default:
			fmt.Println(string(msg), "with default")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
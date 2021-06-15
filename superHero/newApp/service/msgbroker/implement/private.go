package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

func (impl implementation) newHandler(topic msgbrokerin.TopicMsgBroker) (handler common.Handler) {
	return func(ctx context.Context, msg []byte) {
		var err error
		switch topic {
		case msgbrokerin.TopicResponse:
			err = impl.usrService.MsgReceiver(ctx, msg)
		default:
			fmt.Println(string(msg), "with default")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
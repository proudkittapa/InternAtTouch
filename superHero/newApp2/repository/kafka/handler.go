package kafka

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"

	//"git.touchdevops.com/touchtechnologies/authentication-service/service/msgbroker/msgbrokerin"
)

func (message Kafka) RegisterHandler(topic msgbrokerin.TopicMsgBroker, handler common.Handler) {
	message.Broker.RegisterHandler(string(topic), handler)
}

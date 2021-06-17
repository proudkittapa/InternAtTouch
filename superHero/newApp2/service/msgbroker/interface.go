package msgbroker

import "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"

type Service interface {
	Receiver(topics []msgbrokerin.TopicMsgBroker)
	Sender(topic msgbrokerin.TopicMsgBroker, data interface{}) (err error)
}
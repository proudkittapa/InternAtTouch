package implement

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"log"
)

func (impl implementation)Receiver(topics []msgbrokerin.TopicMsgBroker)(){
	for _, topic := range topics{
		log.Println("receivers",topic)
		impl.msgBroker.RegisterHandler(topic, impl.newHandler(topic))
	}
	impl.msgBroker.Consumer()
}
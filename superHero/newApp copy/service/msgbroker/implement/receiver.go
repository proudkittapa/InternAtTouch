package implement

import (
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"log"
)

func (impl implementation)Receiver(topics []msgbrokerin.TopicMsgBroker)(){
	fmt.Println("enter receiver")
	for _, topic := range topics{
		log.Println(topic)
		log.Println(len(topics))
		impl.msgBroker.RegisterHandler(topic, impl.newHandler(topic))
		impl.msgBroker.Consumer()
	}

}
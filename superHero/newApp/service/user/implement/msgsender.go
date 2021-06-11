package implement

import (
	"encoding/json"
	"errors"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

const InvalidInputTypeErr string = "invalid authentication type"

func (impl *implementation) MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {

	switch topic {
	case msgbrokerin.TopicUser:
		err = impl.sender(topic, input)
		if err != nil {
			return err
		}
	}
	return
}

func (impl *implementation) sender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	create, ok := input.(userin.MsgBrokerCreate) //set data that will be send to kafka
	if !ok {
		return errors.New(InvalidInputTypeErr)
	}

	msg, err := json.Marshal(create)
	if err != nil {
		return err
	}

	err = impl.mBroker.Producer(topic, msg)
	if err != nil {
		return err
	}

	return
}

//func (impl *implementation) senderRequestPassword(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
//	create, ok := input.(out.MsgBroker)
//	if !ok {
//		return errors.New(InvalidInputTypeErr)
//	}
//
//	msg, err := json.Marshal(create)
//	if err != nil {
//		return err
//	}
//
//	err = impl.mBroker.Producer(topic, msg)
//	if err != nil {
//		return err
//	}
//
//	return
//}
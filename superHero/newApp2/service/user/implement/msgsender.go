package implement

import (
	"encoding/json"
	"errors"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user/userin"
)

const InvalidInputTypeErr string = "invalid authentication type"

func (impl *implementation) MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {

	switch topic {
	case msgbrokerin.TopicResponseCreate:
		err = impl.senderResponseCreate(topic, input)
		if err != nil {
			return err
		}
	case msgbrokerin.TopicResponseUpdate:
		err = impl.senderResponseUpdate(topic, input)
		if err != nil {
			return err
		}
	case msgbrokerin.TopicResponseDelete:
		err = impl.senderResponseDelete(topic, input)
		if err != nil {
			return err
		}

	}

	return
}

func (impl *implementation) senderResponseCreate(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
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

func (impl *implementation) senderResponseUpdate(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	create, ok := input.(userin.MsgBrokerUpdate) //set data that will be send to kafka
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
func (impl *implementation) senderResponseDelete(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	create, ok := input.(userin.MsgBrokerDelete) //set data that will be send to kafka
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

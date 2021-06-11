package kafka

import "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"

//import "git.touchdevops.com/touchtechnologies/authentication-service/service/msgbroker/msgbrokerin"

func (message Kafka) Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error) {
	err = message.SendTopicMessage(string(topic), msg)
	if err != nil {
		return err
	}

	return nil
}
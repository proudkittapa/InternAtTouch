package implement

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user"
)

type implementation struct{
	msgBroker util.RepositoryMsgBroker
	usrService userService.Service
}

func New(msgBroker util.RepositoryMsgBroker, usrService userService.Service) (service msgbroker.Service){
	impl := &implementation{
		msgBroker: msgbroker,
		usrService: usrService,
	}
	return impl
}
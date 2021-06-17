package implement

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker"
	userService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/util"

)

type implementation struct{
	msgBroker util.RepositoryMsgBroker
	usrService userService.Service
}

func New(msgBroker util.RepositoryMsgBroker, usrService userService.Service) (service msgbroker.Service){
	impl := &implementation{
		msgBroker: msgBroker,
		usrService: usrService,
	}
	return impl
}
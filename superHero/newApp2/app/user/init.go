package user

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) (ctrl *Controller) {
	return &Controller{service}
}

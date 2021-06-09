package user

import (
	"touch/service/user"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) (ctrl *Controller) {
	return &Controller{service}
}

package implement

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"
	"touch/service/user"
	"touch/service/validator"
)

type implementation struct {
	validator validator.Validator
}

func New(validator validator.Validator) (service user.Service) {
	return &implementation{validator}
}

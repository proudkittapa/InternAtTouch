package implement

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/util"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	mBroker util.RepositoryMsgBroker
}

func New(validator validator.Validator, repo util.Repository, mBroker util.RepositoryMsgBroker) (service user.Service) {
	return &implementation{validator, repo, mBroker}
}

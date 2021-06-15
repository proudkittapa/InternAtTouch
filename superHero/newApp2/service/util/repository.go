package util

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type Repository interface {
	Create(ctx context.Context, ent interface{}) (err error)
	Update(ctx context.Context, ent interface{}, ID string) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	View(ctx context.Context, id string) (a domain.InsertQ, err error)
	ViewAll(ctx context.Context, perPage int, page int) (a []domain.InsertQ, err error)
	CheckExistName(ctx context.Context, name string) (bool, error)
	CheckExistActualName(ctx context.Context, actualName string) (bool, error)
}

type RepositoryUsers interface{
	Repository
}
type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}
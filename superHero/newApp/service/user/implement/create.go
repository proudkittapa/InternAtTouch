package implement

import (
	"context"
	"fmt"
	"log"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/modern-go/reflect2"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	defer func() {
		if !reflect2.IsNil(err) {
			return
		}
		if err == impl.sendMsgCreate(input) {
			log.Println(err)
		}

	}()
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validte", err)
		return "validate error", err
	}

	//user := userin.CreateInputToUserDomain(input)
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	//err = impl.repo.Create(ctx, user)
	// fmt.Println("output create:", user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (impl *implementation) sendMsgCreate(input *userin.CreateInput) (err error) {
	return impl.MsgSender(msgbrokerin.TopicUser, userin.MsgBrokerCreate{
		Action:     msgbrokerin.ActionCreate,
		ID:         input.ID,
		Name:       input.Name,
		ActualName: input.ActualName,
		Gender:     input.Gender,
		BirthDate:  input.BirthDate,
		Height:     input.Height,
		SuperPower: input.SuperPower,
		Alive:      input.Alive,
	})
}

package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/modern-go/reflect2"
	"log"

	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		if err == impl.sendMsgUpdate(input){
			log.Println(err)
		}
	}()
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validte", err)
		return "validate error", err
	}

	user := userin.UpdateInputToUserDomain(input)

	//err = impl.repo.Update(ctx, user, user.ID)
	if err != nil {
		// fmt.Println("er")
		return "", err
	}

	return user.Name, nil
}

func (impl *implementation) sendMsgUpdate(input *userin.UpdateInput) (err error) {
	return impl.MsgSender("update", userin.MsgBrokerCreate{
		Action:     msgbrokerin.ActionUpdate,
		ID:             input.ID,
		Name:           input.Name,
		ActualName:     input.ActualName,
		ActualLastName: input.ActualLastName,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		Height:         input.Height,
		SuperPower:     input.SuperPower,
		Alive:          input.Alive,
		Universe:       input.Universe,
		Movies:         input.Movies,
		Enemies:        input.Enemies,
		FamilyMember:   input.FamilyMember,
		About:          input.About,
	})
}

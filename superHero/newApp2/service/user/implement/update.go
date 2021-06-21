package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user/userin"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
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
	return impl.MsgSender("responseUpdate", userin.MsgBrokerUpdate{
		Action:     msgbrokerin.ActionUpdateResponse,
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
		Code: input.Code,
		Err: input.Err,
	})
}
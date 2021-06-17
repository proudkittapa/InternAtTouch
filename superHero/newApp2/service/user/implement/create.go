package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user/userin"
	"github.com/modern-go/reflect2"
	"log"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	//var msg []byte
	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		fmt.Println("sendMsgCreate, response")
		if err == impl.sendMsgCreate(input){
			log.Println(err)
		}
		//if err == impl.receiverMsgCreate(){
		//	log.Println(err)
		//}
		//err = impl.MsgReceiver(ctx, msg)
		//if err != nil{
		//	log.Println(err)
		//}
	}()
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
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
	return impl.MsgSender("responseCreate", userin.MsgBrokerCreate{
		Action:     msgbrokerin.ActionCreateResponse,
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

//func (impl *implementation) receiverMsgCreate(msg []byte) (err error) {
//	ctx := context.Background()
//	err = impl.MsgReceiver(ctx, msg )
//	return err
//}
package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/service/user/userin"
	"github.com/modern-go/reflect2"
	//"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/util"
	"log"
)

func (impl *implementation) MsgReceiver(ctx context.Context, msg []byte) (err error) {
	msgInput := &userin.MsgBrokerCreate{}
	err = json.Unmarshal(msg, msgInput)
	if err != nil {
		return err
	}

	fmt.Println(string(msg))
	switch msgInput.Action {
	case msgbrokerin.ActionCreate:
		fmt.Println("receive action response")
		err = impl.receiveCreateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionUpdate:
		fmt.Println("receive update response")
		err = impl.receiveUpdateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionDelete:
		fmt.Println("receive delete response")
		err = impl.receiveDeleteAction(ctx, msgInput)
		if err != nil {
			return err
		}

	}

	return
}


func (impl *implementation) receiveCreateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {

	input := msgBrokerInput.ToCreateInput()
	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		fmt.Println("response create")
		if err == impl.sendMsgCreate(input){
			log.Println(err)
		}
	}()
	domainUser := input.CreateInputToUserDomain()
	fmt.Println("reached receive create action")
	//err = impl.repo.Create(ctx, domainUser)
	err = impl.validator.Validate(input)
	fmt.Println("err validate", err)
	if err!= nil{
		input.Code = 422
		input.Err = err
		impl.sendMsgCreate(input)
		return err
	}
	err = impl.repo.Insert(ctx, domainUser)
	input.Err = err
	fmt.Println("err app2 insert", input.Err)
	if err == nil{
		input.Code = 200
	}else{
		input.Code = 422
	}


	return nil
}

func (impl *implementation) receiveUpdateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToUpdateInput()
	domainUser := input.UpdateInputToUserDomain()
	fmt.Println("reached receive Update action")
	err = impl.validator.Validate(input)
	if err!= nil{
		input.Code = 422
		input.Err = err
		return err
	}
	err = impl.repo.Update(ctx, domainUser)
	input.Err = err
	//fmt.Println()
	if err == nil{
		input.Code = 200
	}else{
		input.Code = 422
	}

	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		fmt.Println("response Update")
		if err == impl.sendMsgUpdate(input){
			log.Println(err)
		}
	}()

	return nil
}
func (impl *implementation) receiveDeleteAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToDeleteInput()
	domainUser := input.DeleteInputToUserDomain()
	//fmt.Println("reached receive create action")
	err = impl.repo.Delete(ctx, domainUser.ID)
	input.Err = err
	if err == nil{
		input.Code = 200
	}else{
		input.Code = 422
	}
	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		fmt.Println("response delete")
		if err == impl.sendMsgDelete(input){
			log.Println(err)
		}
	}()

	return nil
}
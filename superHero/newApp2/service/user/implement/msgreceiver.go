package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
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
	domainUser := input.CreateInputToUserDomain()
	fmt.Println("reached receive create action")
	//err = impl.repo.Create(ctx, domainUser)
	err = impl.repo.Insert(ctx, domainUser)
	defer func(){
		if !reflect2.IsNil(err){
			return
		}
		fmt.Println("response create")
		if err == impl.sendMsgCreate(input){
			log.Println(err)
		}
	}()

	return nil
}

func (impl *implementation) receiveUpdateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToUpdateInput()
	domainUser := input.UpdateInputToUserDomain()
	fmt.Println("reached receive Update action")
	err = impl.repo.Update(ctx, domainUser)
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
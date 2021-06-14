package userin

import "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"

type MsgBrokerCreate struct{
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID         string   `json:"id"`
	Name       string   `json:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
}

func (msg MsgBrokerCreate) ToCreateInput()(createInput *CreateInput){
	createInput = &CreateInput{
		ID:         msg.ID,
		Name:       msg.Name,
		ActualName: msg.ActualName,
		Gender:     msg.Gender,
		BirthDate:  msg.BirthDate,
		Height:     msg.Height,
		SuperPower: msg.SuperPower,
		Alive:      msg.Alive,
	}
	return createInput
}


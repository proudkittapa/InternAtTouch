package msgbrokerin

type ActionMsgBroker string

const(
	ActionUpsert ActionMsgBroker = "create"
	//ActionUpdate ActionMsgBroker = "update"
	ActionDelete ActionMsgBroker = "delete"
)
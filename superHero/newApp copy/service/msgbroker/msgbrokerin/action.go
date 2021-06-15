package msgbrokerin

type ActionMsgBroker string

const(
	ActionCreate ActionMsgBroker = "create"
	//ActionResponse ActionMsgBroker = "response"
	ActionUpdate ActionMsgBroker = "update"
	ActionDelete ActionMsgBroker = "delete"
)
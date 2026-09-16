package model

type ActionState struct {
	ActionID          string `json:"actionId"`
	ActionType        string `json:"actionType"`
	ActionDescription string `json:"actionDescription"`
	ActionStatus      string `json:"actionStatus"`
	ResultDescription string `json:"resultDescription"`
}

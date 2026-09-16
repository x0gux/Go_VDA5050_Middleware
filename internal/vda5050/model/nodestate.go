package model

type NodeState struct {
	NodeID          string `json:"nodeId"`
	SequenceID      int    `json:"sequenceId"`
	NodeDescription string `json:"nodeDescription,omitempty"`
	Released        bool   `json:"released"`
}

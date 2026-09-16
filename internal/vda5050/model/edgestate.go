package model

type EdgeState struct {
	EdgeID          string `json:"edgeId"`
	SequenceID      int    `json:"sequenceId"`
	EdgeDescription string `json:"edgeDescription,omitempty"`
	Released        bool   `json:"released"`
}

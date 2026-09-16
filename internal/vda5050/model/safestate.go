package model

type SafetyState struct {
	EStop          string `json:"eStop"` // NONE, MANUAL, REMOTE, AUTOACK
	FieldViolation bool   `json:"fieldViolation"`
}

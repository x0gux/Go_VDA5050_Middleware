package model

type VdaError struct {
	ErrorType        string     `json:"errorType"`
	ErrorReferences  []ErrorRef `json:"errorReferences"`
	ErrorDescription string     `json:"errorDescription"`
	ErrorLevel       string     `json:"errorLevel"` // WARNING, FATAL
}

package model

type VdaInfo struct {
	InfoType        string     `json:"infoType"`
	InfoReferences  []ErrorRef `json:"infoReferences"`
	InfoDescription string     `json:"infoDescription"`
	InfoLevel       string     `json:"infoLevel"` // INFO, DEBUG
}

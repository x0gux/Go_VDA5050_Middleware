package model

type CustomInfo struct {
	FullChargeRequest bool   `json:"fullChargeRequest"`
	Host              string `json:"host"` // 로봇이 스스로 보고하는 자신의 IP (읽기 전용)
	NaviType          int    `json:"naviType"`
	RobotType         string `json:"robotType"`
}

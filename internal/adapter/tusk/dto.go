package tusk

// Tusk(VDA5050) 원본 Raw Payload 파싱용 구조체
type State struct {
	SerialNumber  string       `json:"serialNumber"`
	Manufacturer  string       `json:"manufacturer"`
	OperatingMode string       `json:"operatingMode"`
	Driving       bool         `json:"driving"`
	BatteryState  BatteryState `json:"batteryState"`
	AgvPosition   AgvPosition  `json:"agvPosition"`
	Custom        CustomInfo   `json:"custom"`
}

type BatteryState struct {
	BatteryCharge float64 `json:"batteryCharge"`
	Charging      bool    `json:"charging"`
}

type AgvPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CustomInfo struct {
	RobotType string `json:"robotType"`
}

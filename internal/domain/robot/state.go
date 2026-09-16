package robot

type State struct {
	SerialNumber string `json:"serial_number"`
	Manufacturer string `json:"manufacturer"`
	RobotType    string `json:"robot_type"`

	OperatingMode string `json:"operating_mode"`
}

type BatteryState struct {
	BatteryCharge float64 `json:"battery_charge"`
	Charging      bool    `json:"charging"`
}

type AgvPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

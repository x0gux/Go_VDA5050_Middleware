package robot

type State struct {
	SerialNumber  string
	Manufacturer  string
	RobotType     string
	OperatingMode string

	BatteryState BatteryState
	AgvPosition  AgvPosition

	Driving bool
}

type BatteryState struct {
	BatteryCharge float64
	Charging      bool
}

type AgvPosition struct {
	X float64
	Y float64
}

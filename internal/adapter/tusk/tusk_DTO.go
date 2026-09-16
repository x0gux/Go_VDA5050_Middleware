package tusk

import "example.com/m/v2/internal/domain"

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

func (s *State) ToDomain() domain.RobotTypeDTO {
	return domain.RobotTypeDTO{
		SerialNumber:  s.SerialNumber,
		Manufacturer:  s.Manufacturer,
		RobotType:     s.Custom.RobotType,
		OperatingMode: s.OperatingMode,
		BatteryCharge: s.BatteryState.BatteryCharge,
		Charging:      s.BatteryState.Charging,
		X:             s.AgvPosition.X,
		Y:             s.AgvPosition.Y,
		Driving:       s.Driving,
	}
}

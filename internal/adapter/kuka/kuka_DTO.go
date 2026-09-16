package kuka

import "example.com/m/v2/internal/domain"

type State struct {
	HeaderID      int          `json:"headerId"`
	Timestamp     string       `json:"timestamp"`
	Manufacturer  string       `json:"manufacturer"`
	SerialNumber  string       `json:"serialNumber"`
	OperatingMode string       `json:"operatingMode"`
	Driving       bool         `json:"driving"`
	AgvPosition   AgvPosition  `json:"agvPosition"`
	BatteryState  BatteryState `json:"batteryState"`
	Information   []Info       `json:"information"`
}

type AgvPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type BatteryState struct {
	BatteryCharge float64 `json:"batteryCharge"`
	Charging      bool    `json:"charging"`
}

type Info struct {
	InfoType       string          `json:"infoType"`
	InfoReferences []InfoReference `json:"infoReferences"`
}

type InfoReference struct {
	ReferenceKey   string `json:"referenceKey"`
	ReferenceValue string `json:"referenceValue"`
}

func (s *State) ToDomain() domain.RobotTypeDTO {

	robotType := "KUKA-AGV"
	for _, info := range s.Information {
		for _, ref := range info.InfoReferences {
			if ref.ReferenceKey == "agvModelName" {
				robotType = ref.ReferenceValue
				break
			}
		}
	}

	return domain.RobotTypeDTO{
		SerialNumber:  s.SerialNumber,
		Manufacturer:  s.Manufacturer,
		RobotType:     robotType,
		OperatingMode: s.OperatingMode,
		BatteryCharge: s.BatteryState.BatteryCharge,
		Charging:      s.BatteryState.Charging,
		X:             s.AgvPosition.X,
		Y:             s.AgvPosition.Y,
		Driving:       s.Driving,
	}
}

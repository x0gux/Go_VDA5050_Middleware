package domain

// RobotTypeDTO : 화면 출력 및 미들웨어 내부 처리를 위한 경량화 DTO
type RobotTypeDTO struct {
	SerialNumber  string  `json:"serialNumber"`
	Manufacturer  string  `json:"manufacturer"`
	RobotType     string  `json:"robotType"`
	OperatingMode string  `json:"operatingMode"`
	BatteryCharge float64 `json:"batteryCharge"`
	Charging      bool    `json:"charging"`
	X             float64 `json:"x"`
	Y             float64 `json:"y"`
	Driving       bool    `json:"driving"`
}

// ToRobotTypeDTO : VDA5050 원본 State에서 필요한 정보만 DTO로 추출
func (s *State) ToRobotTypeDTO() RobotTypeDTO {
	return RobotTypeDTO{
		SerialNumber:  s.SerialNumber,
		Manufacturer:  s.Manufacturer,
		RobotType:     s.Custom.RobotType, // CustomInfo 확장 필드에서 추출
		OperatingMode: s.OperatingMode,
		BatteryCharge: s.BatteryState.BatteryCharge,
		Charging:      s.BatteryState.Charging,
		X:             s.AgvPosition.X,
		Y:             s.AgvPosition.Y,
		Driving:       s.Driving,
	}
}

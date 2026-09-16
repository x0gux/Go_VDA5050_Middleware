package model

type BatteryState struct {
	BatteryCharge  float64 `json:"batteryCharge"`  // %
	BatteryVoltage float64 `json:"batteryVoltage"` // mV (제조사에 따라 단위 상이할 수 있음, 확인 필요)
	BatteryHealth  float64 `json:"batteryHealth,omitempty"`
	Charging       bool    `json:"charging"`
	Reach          float64 `json:"reach,omitempty"`
}

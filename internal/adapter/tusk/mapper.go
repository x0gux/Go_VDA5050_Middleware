package tusk

func (s *State) ToDomain() State {
	return State{
		SerialNumber:  s.SerialNumber,
		Manufacturer:  s.Manufacturer,
		OperatingMode: s.OperatingMode,
		BatteryState: BatteryState{
			BatteryCharge: s.BatteryState.BatteryCharge,
			Charging:      s.BatteryState.Charging,
		},
		AgvPosition: AgvPosition{
			X: s.AgvPosition.X,
			Y: s.AgvPosition.Y,
		},
		Driving: s.Driving,
	}
}

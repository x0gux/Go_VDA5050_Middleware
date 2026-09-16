package tusk

import "example.com/m/v2/internal/domain/robot"

func (s *State) ToDomain() robot.State {
	return robot.State{
		SerialNumber:  s.SerialNumber,
		Manufacturer:  s.Manufacturer,
		RobotType:     s.Custom.RobotType,
		OperatingMode: s.OperatingMode,
		BatteryState: robot.BatteryState{
			BatteryCharge: s.BatteryState.BatteryCharge,
			Charging:      s.BatteryState.Charging,
		},
		AgvPosition: robot.AgvPosition{
			X: s.AgvPosition.X,
			Y: s.AgvPosition.Y,
		},
		Driving: s.Driving,
	}
}

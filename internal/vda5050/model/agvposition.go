package model

type AgvPosition struct {
	PositionInitialized bool    `json:"positionInitialized"`
	LocalizationScore   float64 `json:"localizationScore"`
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"`
}

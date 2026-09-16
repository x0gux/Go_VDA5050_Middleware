package model

type MapInfo struct {
	MapID      string `json:"mapId"`
	MapVersion string `json:"mapVersion"`
	MapStatus  string `json:"mapStatus"` // ENABLED, DISABLED
}

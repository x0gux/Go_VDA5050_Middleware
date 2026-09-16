package model

type Load struct {
	LoadID         string  `json:"loadId,omitempty"`
	LoadType       string  `json:"loadType,omitempty"`
	LoadPosition   string  `json:"loadPosition,omitempty"`
	BoundingBoxRef any     `json:"boundingBoxReference,omitempty"`
	LoadDimensions any     `json:"loadDimensions,omitempty"`
	Weight         float64 `json:"weight,omitempty"`
}

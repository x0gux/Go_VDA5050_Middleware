package model

type State struct {
	// ---- 헤더 / 식별 정보 ----
	HeaderID     int64  `json:"headerId"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	Manufacturer string `json:"manufacturer"`
	SerialNumber string `json:"serialNumber"`

	// ---- 주문(Order) 진행 상태 ----
	OrderID            string `json:"orderId"`
	OrderUpdateID      int    `json:"orderUpdateId"`
	LastNodeID         string `json:"lastNodeId"`
	LastNodeSequenceID int    `json:"lastNodeSequenceId"`

	NodeStates []NodeState `json:"nodeStates"`
	EdgeStates []EdgeState `json:"edgeStates"`

	// ---- 액션 실행 상태 ----
	ActionStates []ActionState `json:"actionStates"`

	// ---- 위치/속도 ----
	AgvPosition AgvPosition `json:"agvPosition"`
	Velocity    Velocity    `json:"velocity"`

	// ---- 맵 정보 ----
	Maps []MapInfo `json:"maps"`

	// ---- 로드(적재물) ----
	Loads []Load `json:"loads"`

	// ---- 동작 상태 ----
	Driving       bool   `json:"driving"`
	Paused        bool   `json:"paused"`
	OperatingMode string `json:"operatingMode"` // e.g. AUTOMATIC, MANUAL, TEACHIN...

	// ---- 배터리 ----
	BatteryState BatteryState `json:"batteryState"`

	// ---- 에러 / 정보 / 안전 ----
	Errors      []VdaError  `json:"errors"`
	Information []VdaInfo   `json:"information"`
	SafetyState SafetyState `json:"safetyState"`

	// ---- 제조사 커스텀 확장 필드 ----
	// host(IP) 등은 표준 필드가 아니라 이 안에만 존재함 (읽기 전용 보고값)
	Custom CustomInfo `json:"custom"`
}

package domain

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

// NodeState : 아직 지나지 않은/현재 처리 중인 노드 상태
type NodeState struct {
	NodeID          string `json:"nodeId"`
	SequenceID      int    `json:"sequenceId"`
	NodeDescription string `json:"nodeDescription,omitempty"`
	Released        bool   `json:"released"`
}

// EdgeState : 아직 지나지 않은/현재 처리 중인 엣지 상태
type EdgeState struct {
	EdgeID          string `json:"edgeId"`
	SequenceID      int    `json:"sequenceId"`
	EdgeDescription string `json:"edgeDescription,omitempty"`
	Released        bool   `json:"released"`
}

// ActionState : 개별 액션(instantAction/orderAction)의 실행 결과
type ActionState struct {
	ActionID          string `json:"actionId"`
	ActionType        string `json:"actionType"`
	ActionDescription string `json:"actionDescription"`
	// WAITING, INITIALIZING, RUNNING, PAUSED, FINISHED, FAILED
	ActionStatus      string `json:"actionStatus"`
	ResultDescription string `json:"resultDescription"`
}

// AgvPosition : 로봇의 현재 위치 및 위치추정(localization) 신뢰도
type AgvPosition struct {
	PositionInitialized bool    `json:"positionInitialized"`
	LocalizationScore   float64 `json:"localizationScore"`
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"`
}

// Velocity : 로봇의 현재 속도
type Velocity struct {
	Vx    float64 `json:"vx"`
	Vy    float64 `json:"vy"`
	Omega float64 `json:"omega"`
}

// MapInfo : 로봇에 탑재된 맵 목록
type MapInfo struct {
	MapID      string `json:"mapId"`
	MapVersion string `json:"mapVersion"`
	MapStatus  string `json:"mapStatus"` // ENABLED, DISABLED
}

// Load : 로봇이 현재 적재 중인 화물 정보 (VDA5050 표준 필드, 비어있을 수 있음)
type Load struct {
	LoadID         string  `json:"loadId,omitempty"`
	LoadType       string  `json:"loadType,omitempty"`
	LoadPosition   string  `json:"loadPosition,omitempty"`
	BoundingBoxRef any     `json:"boundingBoxReference,omitempty"`
	LoadDimensions any     `json:"loadDimensions,omitempty"`
	Weight         float64 `json:"weight,omitempty"`
}

// BatteryState : 배터리 상태
type BatteryState struct {
	BatteryCharge  float64 `json:"batteryCharge"`  // %
	BatteryVoltage float64 `json:"batteryVoltage"` // mV (제조사에 따라 단위 상이할 수 있음, 확인 필요)
	BatteryHealth  float64 `json:"batteryHealth,omitempty"`
	Charging       bool    `json:"charging"`
	Reach          float64 `json:"reach,omitempty"`
}

// VdaError : 에러 정보
type VdaError struct {
	ErrorType        string     `json:"errorType"`
	ErrorReferences  []ErrorRef `json:"errorReferences"`
	ErrorDescription string     `json:"errorDescription"`
	ErrorLevel       string     `json:"errorLevel"` // WARNING, FATAL
}

// ErrorRef : 에러에 딸린 참조 키/값 (orderId, nodeId 등)
type ErrorRef struct {
	ReferenceKey   string `json:"referenceKey"`
	ReferenceValue string `json:"referenceValue"`
}

// VdaInfo : 정보성 메시지 (에러는 아니지만 참고할 정보)
type VdaInfo struct {
	InfoType        string     `json:"infoType"`
	InfoReferences  []ErrorRef `json:"infoReferences"`
	InfoDescription string     `json:"infoDescription"`
	InfoLevel       string     `json:"infoLevel"` // INFO, DEBUG
}

// SafetyState : 안전 관련 상태
type SafetyState struct {
	EStop          string `json:"eStop"` // NONE, MANUAL, REMOTE, AUTOACK
	FieldViolation bool   `json:"fieldViolation"`
}

// CustomInfo : 제조사(Tuskrobots) 확장 필드
// 표준 VDA5050 필드가 아니므로 서버->로봇 명령에는 사용 불가, 참고/모니터링 용도로만 사용
type CustomInfo struct {
	FullChargeRequest bool   `json:"fullChargeRequest"`
	Host              string `json:"host"` // 로봇이 스스로 보고하는 자신의 IP (읽기 전용)
	NaviType          int    `json:"naviType"`
	RobotType         string `json:"robotType"`
}

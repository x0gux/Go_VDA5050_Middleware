# Go VDA5050 Middleware

Go 기반의 VDA5050 Middleware 프로젝트입니다.

서로 다른 AMR 제조사의 로봇 인터페이스를 Middleware에서 추상화하고, 공통 Domain 모델과 Repository를 통해 로봇 상태를 관리하는 것을 목표로 합니다.

## Architecture

```text
MQTT Broker
    ↓
Subscriber
    ↓
Router
    ↓
Vendor Adapter
    ↓
Mapper
    ↓
Domain Robot State
    ↓
Repository
```

## Project Structure

```text
.
├── cmd/
│   └── middleware.go
│
├── internal/
│   ├── adapter/
│   │   ├── kuka/
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── error.go
│   │   │
│   │   └── tusk/
│   │       ├── dto.go
│   │       ├── mapper.go
│   │       └── error.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── domain/
│   │   └── robot/
│   │       ├── state.go
│   │       └── error.go
│   │
│   ├── repository/
│   │   ├── robot_store.go
│   │   └── memory/
│   │       └── robot_store.go
│   │
│   ├── service/
│   │   └── error.go
│   │   └── robot_service.go
│   │
│   ├── transport/
│   │   └── mqtt/
│   │       ├── client.go
│   │       ├── mqtt_type.go
│   │       ├── publisher.go
│   │       ├── router.go
│   │       ├── subscriber.go
│   │       └── error.go
│   │
│   └── vda5050/
│       └── model/
│           ├── action.go
│           ├── actionstate.go
│           ├── agvposition.go
│           ├── batterystate.go
│           ├── connection.go
│           ├── custominfo.go
│           ├── edge.go
│           ├── edgestate.go
│           ├── error.go
│           ├── header.go
│           ├── information.go
│           ├── load.go
│           ├── node.go
│           ├── nodestate.go
│           ├── order.go
│           ├── safety_state.go
│           ├── state.go
│           └── velocity.go
│
├── .env
├── .gitignore
├── go.mod
└── README.md
```

## Package Responsibilities

### `cmd`

Middleware 애플리케이션의 시작점입니다.

담당:
- 환경변수 로드
- Repository 생성
- MQTT Client 생성
- Router 생성
- Vendor Handler 등록
- MQTT Subscribe
- Middleware 종료 처리

### `internal/adapter`

Vendor별 DTO와 변환 로직을 관리합니다.

```text
TUSK JSON
   ↓
adapter/tusk.State
   ↓
ToDomain()
   ↓
domain/robot.State
```

KUKA 역시 동일한 방식으로 처리합니다.

> DTO는 Vendor Adapter 내부에 남기고 Domain에는 Vendor 의존성을 넣지 않습니다.

### `internal/domain`

Middleware 내부에서 사용하는 공통 비즈니스 모델입니다.

예:

```go
type State struct {
    SerialNumber  string
    Manufacturer  string
    RobotType     string
    OperatingMode string

    BatteryState BatteryState
    AgvPosition  AgvPosition

    Driving bool
}
```

Domain은 TUSK/KUKA의 JSON 구조를 알지 않아야 합니다.

### `internal/repository`

Domain 데이터를 저장하고 조회합니다.

```go
type RobotStore interface {
    SaveState(state robot.State)
    GetState(serial string) (robot.State, bool)
    GetAllStates() []robot.State
}
```

현재는 Memory Repository를 사용합니다.

```text
repository/
├── robot_store.go
└── memory/
    └── robot_store.go
```

향후 DB가 필요하면 별도 구현을 추가할 수 있습니다.

### `internal/service`

Business Logic과 Use Case를 담당합니다.

```text
Service
   ↓
RobotStore
```

저장 방식이 Memory인지 DB인지 Service가 알 필요가 없도록 구성합니다.

### `internal/transport/mqtt`

MQTT 통신을 역할별로 분리합니다.

```text
mqtt/
├── client.go
├── mqtt_type.go
├── publisher.go
├── router.go
└── subscriber.go
```

- `client.go`: MQTT Broker 연결 및 Client 관리
- `publisher.go`: MQTT Publish
- `subscriber.go`: MQTT Subscribe 및 메시지 수신
- `router.go`: Topic 분석 및 Handler dispatch
- `mqtt_type.go`: MQTT 관련 공통 Type

## MQTT Message Flow

예를 들어 TUSK Robot State가 들어오는 경우:

```text
TUSK Robot
    ↓
MQTT Broker
    ↓
Subscriber
    ↓
Router
    ↓
TUSK DTO
    ↓
ToDomain()
    ↓
robot.State
    ↓
RobotStore.SaveState()
```

예상 로그:

```text
[MQTT RECEIVE] topic=APR/v2/Tuskrobots/36040/state payload=1234 bytes

[TUSK STATE] robot=36040 battery=82.0% position=(4.143, 5.099) driving=false
```

## Router

Router는 Vendor와 Message Type을 기준으로 Handler를 선택합니다.

```go
router.Register(
    "Tuskrobots",
    "state",
    func(vendor, serial, msgType string, payload []byte) {
        // TUSK State 처리
    },
)
```

KUKA:

```go
router.Register(
    "KUKA",
    "state",
    func(vendor, serial, msgType string, payload []byte) {
        // KUKA State 처리
    },
)
```

Router의 역할은 Business Logic을 처리하는 것이 아니라 메시지를 적절한 Handler로 전달하는 것입니다.

## Vendor Adapter

### TUSK

```text
internal/adapter/tusk/
├── dto.go
└── mapper.go
```

```text
TUSK Message
    ↓
TUSK DTO
    ↓
ToDomain()
    ↓
robot.State
```

### KUKA

```text
internal/adapter/kuka/
├── dto.go
└── mapper.go
```

```text
KUKA Message
    ↓
KUKA DTO
    ↓
ToDomain()
    ↓
robot.State
```

Vendor가 추가되어도 공통 Domain 구조를 유지할 수 있습니다.

## VDA5050 Model

```text
internal/vda5050/model/
```

VDA5050 메시지에 필요한 표준 모델을 관리합니다.

주요 모델:
- State
- Order
- Node
- Edge
- NodeState
- EdgeState
- Action
- ActionState
- BatteryState
- AgvPosition
- Velocity
- Error
- SafetyState
- Load
- Information
- CustomInfo

Vendor Domain Model과 VDA5050 Model은 목적이 다르므로 직접 섞지 않습니다.

## Environment

`.env` 예시:

```env
MQTT_BROKER=tcp://localhost:1883

MQTT_USERNAME=
MQTT_PASSWORD=

MQTT_TUSK_CLIENT_ID=vda5050-tusk-middleware
MQTT_KUKA_CLIENT_ID=vda5050-kuka-middleware
```

실제 MQTT Broker 환경에 맞게 수정합니다.

`.env`에는 인증 정보가 포함될 수 있으므로 Git에 커밋하지 않는 것을 권장합니다.

## Run

Go 버전:

```text
Go 1.27.1
```

의존성 정리:

```bash
go mod tidy
```

실행:

```bash
go run ./cmd
```

Build:

```bash
go build -o middleware ./cmd
```

Windows에서는:

```bash
middleware.exe
```

## Repository 확인

수신한 Robot State는 현재 Memory Repository에 저장됩니다.

```go
state, exists := robotStore.GetState("36040")
```

전체 Robot 조회:

```go
robots := robotStore.GetAllStates()
```

예상 데이터:

```text
Robot ID : 36040
Battery  : 82%
Position : X=4.143 Y=5.099
Driving  : false
```

## Publisher

MQTT Publisher는 Middleware가 Robot 또는 외부 시스템으로 메시지를 전송할 때 사용합니다.

```go
err := client.Publish(
    topic,
    0,
    payload,
)
```

Publisher는 Vendor-specific Business Logic을 처리하지 않습니다.

```text
Business Logic
      ↓
Publisher
      ↓
MQTT Broker
```

## Design Principles

### 1. Vendor DTO와 Domain 분리

잘못된 구조:

```text
TUSK DTO
 ↓
Service
 ↓
KUKA DTO
```

권장:

```text
TUSK DTO ──┐
           ├──> Domain
KUKA DTO ──┘
```

### 2. MQTT와 Vendor Adapter 분리

Vendor Adapter가 MQTT Client를 직접 생성하지 않습니다.

권장:

```text
MQTT Transport
    ↓
Router
    ↓
Vendor Adapter
```

### 3. Repository와 Service 분리

Repository:

```text
데이터 저장 / 조회
```

Service:

```text
Business Logic / Use Case
```

### 4. Domain은 외부 기술을 몰라야 함

Domain에서 MQTT, TUSK API, KUKA API, Paho MQTT 등의 외부 기술에 직접 의존하지 않습니다.

## Development Roadmap

현재:

```text
MQTT Client
    ↓
Publisher / Subscriber
    ↓
Router
    ↓
Vendor DTO
    ↓
Vendor Mapper
    ↓
Domain Robot State
    ↓
Repository / Memory Store
```

다음 단계:

```text
Repository
    ↓
Robot Service
    ↓
VDA5050 State 변환
    ↓
VDA5050 MQTT Publish
```

이후:

```text
Order
 ↓
Order Validation
 ↓
Order / Task Management
 ↓
Vendor별 Command 변환
 ↓
AMR
```

## Goal Architecture

```text
                         ┌──────────────┐
                         │ MQTT Broker  │
                         └──────┬───────┘
                                │
                 ┌──────────────┴──────────────┐
                 │                             │
          ┌──────▼──────┐              ┌──────▼──────┐
          │ Subscriber  │              │  Publisher  │
          └──────┬──────┘              └──────▲──────┘
                 │                            │
                 ▼                            │
          ┌──────────────┐                    │
          │    Router    │                    │
          └──────┬───────┘                    │
                 │                            │
        ┌────────┴────────┐                   │
        ▼                 ▼                   │
   ┌─────────┐       ┌─────────┐             │
   │  TUSK   │       │  KUKA   │             │
   │ Adapter │       │ Adapter │             │
   └────┬────┘       └────┬────┘             │
        │                 │                   │
        └────────┬────────┘                   │
                 ▼                            │
          ┌──────────────┐                    │
          │    Domain    │                    │
          └──────┬───────┘                    │
                 │                            │
          ┌──────▼───────┐                    │
          │   Service    │────────────────────┘
          └──────┬───────┘
                 │
          ┌──────▼───────┐
          │  Repository  │
          └──────────────┘
```

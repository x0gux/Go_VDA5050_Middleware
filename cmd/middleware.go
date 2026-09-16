package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/v2/internal/adapter/kuka"
	"example.com/m/v2/internal/adapter/tusk"
	"example.com/m/v2/internal/repository/memory"
	"example.com/m/v2/internal/transport/mqtt"

	"github.com/joho/godotenv"
)

func main() {
	// --------------------------------------------------
	// 1. 환경변수 로드
	// --------------------------------------------------

	if err := godotenv.Load(); err != nil {
		log.Println(".env 파일을 찾을 수 없습니다. 환경변수를 직접 사용합니다.")
	}

	tusk_broker := os.Getenv("TUSK_BROKER")
	kuka_broker := os.Getenv("KUKA_BROKER")
	tuskClientID := os.Getenv("MQTT_TUSK_CLIENT_ID")
	kukaClientID := os.Getenv("MQTT_KUKA_CLIENT_ID")

	if tusk_broker == "" {
		log.Fatal("TUSK_BROKER가 설정되지 않았습니다.")
	}
	if kuka_broker == "" {
		log.Fatal("KUKA_BROKER가 설정되지 않았습니다.")
	}

	if tuskClientID == "" {
		tuskClientID = "vda5050-tusk-middleware"
	}

	if kukaClientID == "" {
		kukaClientID = "vda5050-kuka-middleware"
	}

	// --------------------------------------------------
	// 2. Repository 생성
	// --------------------------------------------------

	robotStore := memory.NewRobotStore()

	log.Println("Robot Repository initialized")

	// --------------------------------------------------
	// 3. MQTT Router 생성
	// --------------------------------------------------

	router := mqtt.NewRouter()

	// --------------------------------------------------
	// 4. TUSK State Handler
	// --------------------------------------------------

	router.Register(
		"Tuskrobots",
		"state",
		func(vendor, serial, msgType string, payload []byte) {

			var state tusk.State

			if err := json.Unmarshal(payload, &state); err != nil {
				log.Printf(
					"[TUSK] State decode failed | robot=%s | error=%v",
					serial,
					err,
				)
				return
			}

			// Vendor DTO → Domain
			domainState := state.ToDomain()

			// Topic에서 받은 serial이 비어있으면 사용
			if domainState.SerialNumber == "" {
				domainState.SerialNumber = serial
			}

			// Repository 저장
			robotStore.SaveState(domainState)

			log.Printf(
				"[TUSK STATE] robot=%s battery=%.1f%% position=(%.3f, %.3f) driving=%v",
				domainState.SerialNumber,
				domainState.BatteryState.BatteryCharge,
				domainState.AgvPosition.X,
				domainState.AgvPosition.Y,
				domainState.Driving,
			)
		},
	)

	// --------------------------------------------------
	// 5. KUKA State Handler
	// --------------------------------------------------

	router.Register(
		"KUKA",
		"state",
		func(vendor, serial, msgType string, payload []byte) {

			var state kuka.State

			if err := json.Unmarshal(payload, &state); err != nil {
				log.Printf(
					"[KUKA] State decode failed | robot=%s | error=%v",
					serial,
					err,
				)
				return
			}

			// Vendor DTO → Domain
			domainState := state.ToDomain()

			if domainState.SerialNumber == "" {
				domainState.SerialNumber = serial
			}

			// Repository 저장
			robotStore.SaveState(domainState)

			log.Printf(
				"[KUKA STATE] robot=%s battery=%.1f%% position=(%.3f, %.3f) driving=%v",
				domainState.SerialNumber,
				domainState.BatteryState.BatteryCharge,
				domainState.AgvPosition.X,
				domainState.AgvPosition.Y,
				domainState.Driving,
			)
		},
	)

	// --------------------------------------------------
	// 6. TUSK MQTT Client 생성
	// --------------------------------------------------

	tuskClient, err := mqtt.NewClient(
		tusk_broker,
		tuskClientID,
		os.Getenv("MQTT_USERNAME"),
		os.Getenv("MQTT_PASSWORD"),
	)

	if err != nil {
		log.Fatalf("TUSK MQTT 연결 실패: %v", err)
	}

	log.Println("TUSK MQTT connected")

	// --------------------------------------------------
	// 7. KUKA MQTT Client 생성
	// --------------------------------------------------

	kukaClient, err := mqtt.NewClient(
		kuka_broker,
		kukaClientID,
		os.Getenv("MQTT_USERNAME"),
		os.Getenv("MQTT_PASSWORD"),
	)

	if err != nil {
		log.Fatalf("KUKA MQTT 연결 실패: %v", err)
	}

	log.Println("KUKA MQTT connected")

	// --------------------------------------------------
	// 8. TUSK Subscribe
	// --------------------------------------------------

	tuskTopic := "APR/v2/Tuskrobots/+/state"

	if err := tuskClient.Subscribe(
		tuskTopic,
		0,
		func(topic string, payload []byte) {

			log.Printf(
				"[MQTT RECEIVE] topic=%s payload=%d bytes",
				topic,
				len(payload),
			)

			router.Dispatch(topic, payload)
		},
	); err != nil {
		log.Fatalf("TUSK Subscribe 실패: %v", err)
	}

	// --------------------------------------------------
	// 9. KUKA Subscribe
	// --------------------------------------------------

	kukaTopic := "APR/v2/KUKA/+/state"

	if err := kukaClient.Subscribe(
		kukaTopic,
		0,
		func(topic string, payload []byte) {

			log.Printf(
				"[MQTT RECEIVE] topic=%s payload=%d bytes",
				topic,
				len(payload),
			)

			router.Dispatch(topic, payload)
		},
	); err != nil {
		log.Fatalf("KUKA Subscribe 실패: %v", err)
	}

	// --------------------------------------------------
	// 10. Middleware 시작
	// --------------------------------------------------

	log.Println("======================================")
	log.Println(" VDA5050 Middleware Started")
	log.Println("======================================")
	log.Printf("K-Broker T-Broker : %s %s", kuka_broker, tusk_broker)
	log.Printf("TUSK   : %s", tuskTopic)
	log.Printf("KUKA   : %s", kukaTopic)

	// --------------------------------------------------
	// 11. 종료 Signal 대기
	// --------------------------------------------------

	sigCh := make(chan os.Signal, 1)

	signal.Notify(
		sigCh,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-sigCh

	log.Println("Middleware shutting down...")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/v2/pkg"
)

func main() {

	// --------------------------------------------------
	// 1. 핸들러 등록 (벤더 및 메시지 타입별 DTO 바인딩)
	// --------------------------------------------------
	r.Register("Tuskrobots", "state", func(vendor, serial, msgType string, payload []byte) {
		log.Printf("[%s:%s] State 수신 (%d bytes)", vendor, serial, len(payload))
		// var state TuskStateDTO
		// json.Unmarshal(payload, &state)
	})

	r.Register("Kukarobots", "state", func(vendor, serial, msgType string, payload []byte) {
		log.Printf("[%s:%s] Kuka State 처리", vendor, serial)
	})

	r.Register("Kukarobots", "connection", func(vendor, serial, msgType string, payload []byte) {
		log.Printf("[%s:%s] Kuka Connection 상태 업데이트", vendor, serial)
	})

	// --------------------------------------------------
	// 2. 브로커 연결 (독립 인스턴스)
	// --------------------------------------------------

	Kuka_name := os.Getenv("KUKA_NAME")
	Kuka_password := os.Getenv("KUKA_PASSWORD")
	kukaClient, err := pkg.NewClient("mqtt://192.168.0.151:1883", "APR_Kuka_Client", Kuka_name, Kuka_password)
	if err != nil {
		log.Fatal(err)
	}

	// 종료 시그널 대기
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

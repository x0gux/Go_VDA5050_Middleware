package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/v2/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env 파일을 찾을 수 없습니다. 시스템 환경변수를 사용합니다.")
	}
	r := pkg.NewRouter()

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
	Tusk_name := os.Getenv("TUSK_NAME")
	Tusk_password := os.Getenv("TUSK_PASSWORD")
	tuskClient, err := pkg.NewClient("mqtt://192.168.0.170:1883", "APR_Tusk_Client", Tusk_name, Tusk_password)
	if err != nil {
		log.Fatal(err)
	}

	Kuka_name := os.Getenv("KUKA_NAME")
	Kuka_password := os.Getenv("KUKA_PASSWORD")
	kukaClient, err := pkg.NewClient("mqtt://192.168.0.151:1883", "APR_Kuka_Client", Kuka_name, Kuka_password)
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------------------------------
	// 3. 와일드카드 구독 및 라우터 Dispatch 연동
	// --------------------------------------------------
	// go routine 호출 불필요 (Subscribe 내부 토큰으로 동기화 처리)
	if err := tuskClient.Subscribe("APR/v2/Tuskrobots/+/state", 0, r.Dispatch); err != nil {
		log.Println(err)
	}

	if err := kukaClient.Subscribe("APR/v2/Kukarobots/+/state", 0, r.Dispatch); err != nil {
		log.Println(err)
	}
	if err := kukaClient.Subscribe("APR/v2/Kukarobots/+/connection", 0, r.Dispatch); err != nil {
		log.Println(err)
	}

	// 종료 시그널 대기
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

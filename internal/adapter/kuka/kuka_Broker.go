package kuka

import (
	"fmt"
	"log"
	"os"

	"example.com/m/v2/pkg"
	"github.com/joho/godotenv"
)

func kukabroker() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env 파일을 찾을 수 없습니다. 시스템 환경변수를 사용합니다.")
	}

	Kuka_name := os.Getenv("KUKA_NAME")
	Kuka_password := os.Getenv("KUKA_PASSWORD")
	Kuka_address := os.Getenv("KUKA_BROKER")

	kukaClient, err := pkg.NewClient(Kuka_address, "APR_Kuka_Client", Kuka_name, Kuka_password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("kukaClient: %v\n", kukaClient)

}

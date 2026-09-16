package tusk

import (
	"log"
	"os"

	"example.com/m/v2/pkg"
	"github.com/joho/godotenv"
)

func tuskBroker() *pkg.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env 파일을 찾을 수 없습니다. 시스템 환경변수를 사용합니다.")
	}
	Tusk_name := os.Getenv("TUSK_NAME")
	Tusk_password := os.Getenv("TUSK_PASSWORD")
	Tusk_address := os.Getenv("TUSK_BROKER")

	tuskClient, err := pkg.NewClient(Tusk_address, "APR_Tusk_Client", Tusk_name, Tusk_password)
	if err != nil {
		log.Fatal(err)
	}

	return tuskClient

}

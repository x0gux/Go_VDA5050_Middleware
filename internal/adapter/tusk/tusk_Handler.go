package tusk

import (
	"example.com/m/v2/internal/pkg"
)

func GetClient() *pkg.Client {
	return tuskBroker()
}

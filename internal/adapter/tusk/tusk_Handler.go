package tusk

import (
	"example.com/m/v2/pkg"
)

func GetClient() *pkg.Client {
	return tuskBroker()
}

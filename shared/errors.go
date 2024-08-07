package shared

import (
	"log"
)

func Check(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

package model_test

import (
	"log"
)

func valid_test() {
	if !model.valid("DimohaZadira@gmail.com") {
		log.Panic("false negative")
	}
}

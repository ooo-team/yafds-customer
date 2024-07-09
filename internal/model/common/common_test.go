package model

import (
	"log"
	"testing"
	// common "github.com/ooo-team/yafds/internal/model/common"
)

func TestValid(t *testing.T) {
	if !Valid("DimohaZadira@gmail.com") {
		t.Errorf("false negative")
	} else {
		log.Print("zaebis")
	}

	if Valid("Zalupa228") {
		t.Errorf("false negative")
	}
}

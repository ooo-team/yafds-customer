package common

import (
	"fmt"
	"log"
	"testing"
)

func TestValid(t *testing.T) {
	if !Valid("DimohaZadira@gmail.com") {
		t.Errorf("false negative")
	} else {
		log.Print("zaebumba")
	}

	if Valid("Zalupa228") {
		t.Errorf("false positive")
	}
}

func TestLoadEnvVar(t *testing.T) {
	fmt.Println(LoadEnvVar("dbUser"))
}

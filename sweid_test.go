package sweid

import (
	"fmt"
	"testing"
)

func TestValidPIN(t *testing.T) {
	pins := []string{
		"196408233234",
		"19640823-3234",
		"6408233234",
		"6408233234",
	}

	for _, pin := range pins {
		if !ValidPN(pin) {
			t.Errorf("Failed validating %v", pin)
		}
	}
}

func TestInvalidPIN(t *testing.T) {
	pins := []string{
		"196480233234",
		"196412323234",
		"201607110136",
		"196408230135",
		"119640823-3234",
		"640823--3234",
		"640810134",
		"asdf",
		"",
		"1234567890",
	}

	for _, pin := range pins {
		if ValidPN(pin) {
			t.Errorf("Failed validating %v", pin)
		}
	}
}

func TestValidON(t *testing.T) {
	pins := []string{
		"212000-0142",
		"212000-1355",
		"556036-0793",
		"802000-6717",
	}

	for _, pin := range pins {
		if !ValidON(pin) {
			t.Errorf("Failed validating %v", pin)
		}
	}
}

func TestInvalidON(t *testing.T) {
	pins := []string{
		"212000-0144",
		"3234",
		"119680823-3234",
		"640823--3234",
		"640810234",
		"asdf",
		"",
		"1234567890",
		"6408233234",
	}

	for _, pin := range pins {
		if ValidON(pin) {
			t.Errorf("Failed validating %v", pin)
		}
	}
}

func ExampleValidPN() {
	if ValidPN("640823-3234") {
		fmt.Println("Correct Personal Identification Number!")
	} else {
		fmt.Println("Invalid Personal Identification Number!")
	}
}

func ExampleValidON() {
	if ValidON("212000-0142") {
		fmt.Println("Correct Organisational Number!")
	} else {
		fmt.Println("Invalid Organisational Number!")
	}
}

package serial

import "testing"
import "fmt"

var dev string = "/dev/ttyUSB0"
var s *serial = New(dev)

func TestPrintln(t *testing.T) {
	if s == nil {
		t.Error("s == nil")
	}
	s.Println("hahaha")
}

func TestRead(t *testing.T) {
	fmt.Println(s.Readln())
}

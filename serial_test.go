package serial

import "testing"
import "fmt"

func TestPrintln(t *testing.T) {
	Println("hello world")
}

func TestRead(t *testing.T) {
	fmt.Println(Read())
}

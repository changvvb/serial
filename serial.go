package serial

import (
	"bufio"
	"fmt"
	"os"
)

//the device of serial, example "/dev/ttyUSB0"

type serial struct {
	Dev        string
	readBuffer chan ([]byte)
	devFile    *os.File
	scanner    *bufio.Scanner
	buffer     []byte
}

func (s *serial) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(s.devFile, format, a...)
}

func (s *serial) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(s.devFile, a...)
}

func New(dev string) *serial {
	var s serial
	var err error
	s.Dev = dev
	s.readBuffer = make(chan []byte)
	s.devFile, err = os.OpenFile(s.Dev, os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s.scanner = bufio.NewScanner(s.devFile)
	return &s
}

func (s *serial) Scan() bool {
	return s.scanner.Scan()
}

func (s *serial) Bytes() []byte {
	return s.scanner.Bytes()
}

func (s *serial) Text() string {
	return s.scanner.Text()
}

func (s *serial) Read() string {
	return string(<-s.readBuffer)
}

func (s *serial) Readln() string {
	b := <-s.readBuffer
	for l, e := range b {
		if e == '\r' || e == '\n' {
			fmt.Println("len:", l)
			b = b[:l]
		}
	}
	return string(b)
}

func (s *serial) ReadBytes() []byte {
	return <-s.readBuffer
}

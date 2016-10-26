package serial

import (
	"fmt"
	"os"
	"time"
)

//the device of serial, example "/dev/ttyUSB0"
var Dev string
var readBuffer chan ([]byte)
var devFile *os.File

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(devFile, format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(devFile, a...)
}

func init() {
	Dev = "/dev/ttyUSB0"
	readBuffer = make(chan []byte)
	var err error
	devFile, err = os.OpenFile(Dev, os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		b := make([]byte, 1024)
		for {
			n, _ := devFile.Read(b)
			if n != 0 {
				readBuffer <- b
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func Read() string {
	return string(<-readBuffer)
}

func Readln() string {
	b := <-readBuffer
	for l, e := range b {
		if e == '\r' || e == '\n' {
			fmt.Println("len:", l)
			b = b[:l]
		}
	}
	return string(b)
}

func ReadBytes() []byte {
	return <-readBuffer
}

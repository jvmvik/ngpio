package ngpio

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	spec := Nano()
	port, err := spec.FindPort(12)
	if err != nil {
		panic(err)
	}
	port.Output()
	port.High()
	time.Sleep(time.Second * 1)
	port.Low()
	time.Sleep(time.Second * 1)
	port.High()
	time.Sleep(time.Second * 1)
	port.Low()
}

package ngpio

import (
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var wg sync.WaitGroup
	spec := Nano()
	list := []int{7, 11, 13, 15}

	var err error
	var port Port

	for _, i := range list {
		wg.Add(1)
		func(i int) {

			port, err = spec.FindPort(i)
			if err != nil {
				panic(err)
			}

			delay := time.Second * 1
			port.Output()
			port.High()
			time.Sleep(delay)
			port.Low()
			time.Sleep(delay)
			port.High()

			defer wg.Done()
		}(i)

	}

	wg.Wait()
}

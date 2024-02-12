package letsgolang

import (
	"fmt"
	"sync"
	"testing"
)

func PrintRoutine(dataChannel chan string, w *sync.WaitGroup) {
	defer w.Done()
	res := <-dataChannel
	fmt.Println("Data :", res)
}

func InsertRoutine(dataChannel chan string, data string, w *sync.WaitGroup) {
	defer w.Done()
	dataChannel <- data
}

func TestParameterGoroutine(t *testing.T) {
	data := make(chan string)
	defer close(data)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go InsertRoutine(data, "Fazrul", wg)
	go PrintRoutine(data, wg)

	wg.Wait()
}

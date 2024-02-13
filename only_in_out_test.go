package letsgolang

import (
	"fmt"
	"testing"
	"time"
)

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(3 * time.Second)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Fazrul Anugrah Sahi"
}

func OnlyOut(channel <-chan string) {
	result := <-channel
	fmt.Println(result)
}

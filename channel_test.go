package letsgolang

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Fazrul"
		time.Sleep(2 * time.Second)
		fmt.Println("This line...")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

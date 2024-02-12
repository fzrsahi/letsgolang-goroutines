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
		time.Sleep(10 * time.Second)
		channel <- "Fazrul"
		fmt.Println("This line...")
	}()
	data := <-channel
	fmt.Println(data)

}

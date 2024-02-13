package letsgolang

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Fazrul"
	channel <- "Anugrah"
	channel <- "Sahi"

	firstName := <-channel
	middleName := <-channel
	fmt.Println(firstName)
	fmt.Println(middleName)

	channel <- "John"

	john := <-channel

	fmt.Println(john)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

package letsgolang

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke: " + strconv.Itoa(i)
		}
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done..")

}

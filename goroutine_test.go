package letsgolang

import (
	"fmt"
	"testing"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello " + name)
}

func counting(number int) {
	fmt.Println("Count :", number)

}

func TestSayHello(t *testing.T) {
	go sayHello("Fazrul")
	fmt.Println("Running This...")
	time.Sleep(1 * time.Second)
}

func TestGourotineTime(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go counting(i)
	}
	time.Sleep(5 * time.Second)

}

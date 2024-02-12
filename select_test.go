package letsgolang

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func TestSelect(t *testing.T) {
	err := make(chan string)
	randomNumber := make(chan int)
	guess := 2

	defer close(err)
	defer close(randomNumber)
	defer fmt.Println("Your Guess :", guess)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go guessRandomNumber(randomNumber, err, guess, wg)

	select {
	case isErr := <-err:
		fmt.Println(isErr)
	case result := <-randomNumber:
		fmt.Println("Very Well! Right Number : ", result)
	}

	wg.Wait()

}

func guessRandomNumber(c chan int, err chan string, guess int, w *sync.WaitGroup) {
	defer w.Done()
	randomNumber := rand.Intn(5)
	if randomNumber != guess {
		err <- "Wrong!, The Right Number is " + strconv.Itoa(randomNumber)
	} else {
		c <- randomNumber
	}
}

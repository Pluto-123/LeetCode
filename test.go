package main

import (
	"fmt"
	"sync"
)

func main() {

	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func(wait *sync.WaitGroup) {
		i := 1
		for {
			select {
			case <-number:
				if i > 26 {
					wait.Done()
					letter <- true
					return
				}

				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}(&wait)

	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wait.Done()
					number <- true
					return
				}
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}(&wait)
	number <- true
	wait.Wait()

	close(letter)
	close(number)
}

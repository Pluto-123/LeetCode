package main

import (
	"fmt"
	"sync"
)

func main() {

	count := 10
	sum := 100

	wg := sync.WaitGroup{}
	c := make(chan struct{}, count)
	defer close(c)

	for i := 0; i < sum; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-c
		}(i)
	}
	wg.Wait()
}

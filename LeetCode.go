package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)

	go func() {
		// 打印偶数
		for i := 0; i <= 100; i += 2 {
			select {
			case <-c1:
				fmt.Println("偶数线程: ", i)
				c2 <- true
				if i == 100 {
					wg.Done()
					return
				}
			}
		}
	}()
	//
	go func() {
		for i := 1; i <= 101; i += 2 {
			select {
			case <-c2:
				fmt.Println("奇数线程: ", i)
				c1 <- true
				if i == 101 {
					wg.Done()
					return
				}
			}
		}
	}()
	c1 <- true
	wg.Wait()
}

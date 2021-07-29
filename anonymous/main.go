package main

import (
	"fmt"
	"sync"
)

func main() {
	s := sync.WaitGroup{}
	s.Add(1)
	go func(x, y int) {
		defer s.Done()
		fmt.Print(x + y)
	}(10, 20)
	s.Wait()
}

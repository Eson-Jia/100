package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, world.")
}

// InterleavePrintOneChan
// 1.交替打印数字和字母
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func InterleavePrintOneChan() {
	notify := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	printNums := func() {
		defer wg.Done()
		for i := 1; i < 10000; {
			fmt.Print(i)
			fmt.Print(i + 1)

			i += 2
			if _, ok := <-notify; !ok {
				break
			}
		}
	}
	printAB := func() {
		defer wg.Done()
		defer close(notify)
		for i := 'A'; i < 'Z'; {
			notify <- struct{}{}
			fmt.Print(string(i))
			fmt.Print(string(i + 1))
			fmt.Print()
			i += 2
		}
	}
	go printNums()
	go printAB()
	wg.Wait()
}

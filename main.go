package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, world.")
	InterleavePrintWithOneGoroutineQuit()
}

//InterleavePrintWithOneGoroutineQuit
//main goroutine 只会等待一个 goroutine 退出，所以会造成一个打印数字的 goroutine 一直阻塞，无法退出，
//最后导致该 goroutine 泄露，但是如果此时 main goroutine 退出，这个也不算是个问题了。
func InterleavePrintWithOneGoroutineQuit() {
	letter, number := make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i += 1
				fmt.Print(i)
				i += 1
			}
			letter <- true
		}
	}()
	go func() {
		i := 'A'
		for i < 'Z' {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Printf("%c", i)
				i += 1
				fmt.Printf("%c", i)
				i += 1
				number <- true
			}
		}
	}()
	number <- true
	wg.Wait()

}

//InterleavePrintWaitTwoGoroutineQuit
//main 会等待两个 goroutine 退出
func InterleavePrintWaitTwoGoroutineQuit() {
	stringsChan := make(chan bool, 1)
	numsChan := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(numsChan)
		for i := 1; i < 100000; i += 2 {
			if _, ok := <-stringsChan; !ok {
				break
			}
			fmt.Printf("%d%d", i, i+1)
			numsChan <- true
		}
	}()
	go func() {
		defer wg.Done()
		defer close(stringsChan)
		for i := 'A'; i <= 'Y'; i += 2 {
			if _, ok := <-numsChan; !ok {
				break
			}
			fmt.Printf("%c%c", i, i+1)
			stringsChan <- true
		}
	}()
	stringsChan <- true
	wg.Wait()
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

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 10; count++ {
		// 捕获 counter 的值
		value := counter

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
		value++
		counter = value
	}
}

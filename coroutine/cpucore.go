package coroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//定义任务队列
var waitGroup sync.WaitGroup

func calculate(num int) {
	for i := 1; i <= 1000000000; i++ {
		num = num + i
		num = num - i
		num = num * i
		num = num / i
	}
	waitGroup.Done()
}

func CoreTest() {
	start := time.Now()
	fmt.Println(runtime.GOMAXPROCS(0))
	for i := 1; i <= 10; i++ {
		waitGroup.Add(1)
		go calculate(i)
	}
	waitGroup.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}

func CpuTest() {
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumGoroutine())
}

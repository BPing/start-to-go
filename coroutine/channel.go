package coroutine

import (
	"fmt"
	"strconv"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func EchoNoCacheChan() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, x+y)
}

func EchoCacheChan() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//chan T              可以接收和发送类型为 T 的数据
//chan<- float64      只可以用来发送 float64 类型的数据
//<-chan int          只可以用来接收 int 类型的数据

func receive(msg <-chan string) {
	for m := range msg {
		fmt.Println("receive message:", m)
	}
	fmt.Println("receive1 end:")
}

func receive1(msg <-chan string) {
	for m := range msg {
		fmt.Println("receive1 message:", m)
	}
	fmt.Println("receive1 end:")
}

func send(msg chan<- string) {
	for i := 0; i < 10; i++ {
		msg <- "message【" + strconv.Itoa(i) + "】"
		time.Sleep(time.Second)
	}
	//close(msg)
}

func EchoChanRS() {
	msg := make(chan string)
	go send(msg)
	go receive(msg)
	go receive1(msg)

	//time.Sleep(time.Second * 15)
}

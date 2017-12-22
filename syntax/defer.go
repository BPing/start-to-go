package syntax

import "fmt"

func EchoDefer() {
	// 延迟的函数调用被压入一个栈中。当函数返回时，
	// 会按照后进先出的顺序调用被延迟的函数调用。
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("start")
	a()
	fmt.Println(c())
}

func a() {
	i := 0
	// 参数值在定义时已经确认
	defer fmt.Println(i)
	i++
	return
}

func c() (i int) {
	// 可以修改返回值
	defer func() { i++ }()
	return 1
}

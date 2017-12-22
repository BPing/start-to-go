package error

import "fmt"

// https://blog.golang.org/defer-panic-and-recover
// try…catch…finally
// defer, panic, recover

func EchoRecover() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println("Recovered in EchoRecover", err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic("f panic happen")
	fmt.Println("b")
	fmt.Println("f")
}

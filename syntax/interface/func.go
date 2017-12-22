package _interface

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func EchoFunc() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

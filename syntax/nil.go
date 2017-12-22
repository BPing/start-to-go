package syntax

import "fmt"

// 零值是：
//
// 数值类型为 `0`，
// 布尔类型为 `false`，
// 字符串为 `""`（空字符串）。
// 指针（引用）类型为 nil

func EchoNil() {
	var i int
	var f float64
	var b bool
	var s string
	var ptr *int
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, ptr == nil)
}

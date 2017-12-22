package _interface

import (
	"fmt"
	"math"
)

// 它们能做什么比它们是什么更重要

type Abser interface {
	Abs() float64
}

func EchoInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser
	if _, ok := a.(MyFloat); ok {
		fmt.Println("MyFloat")
	}
	a = &v // a *Vertex 实现了 Abser
	if _, ok := a.(*Vertex); ok {
		fmt.Println("*Vertex")
	}

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

package _interface

// https://github.com/luciotato/golang-notes/blob/master/OOP.md

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	r.width = 100
	return 2*r.width + 2*r.height
}

func EchoObj() {
	r := &rect{3, 4}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())
}

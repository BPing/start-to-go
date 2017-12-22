package _interface

import "fmt"

// http://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html
// 内嵌或组合

type base struct {
	a string
	b int
}

type derived struct {
	base // embedding
	d    int
	a    float32 //-SHADOWS
}

func EchoEmbedd() {
	var x derived

	fmt.Printf("%T\n", x.a) //=> x.a, float32 (derived.a 覆盖 base.a)

	fmt.Printf("%T\n", x.base.a) //=> x.base.a, string (访问覆盖成员变量)
}

type NamedObj struct {
	Name string
}

type Shape struct {
	NamedObj  //继承
	color     int32
	isRegular bool
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	NamedObj            //多继承
	Shape               //^^
	center        Point //标准组合
	Width, Height float64
}

func EchoMultiEmbedd() {
	var aRect = Rectangle{
		NamedObj{"NamedObj"},
		Shape{NamedObj{"Shape"}, 0, true},
		Point{0, 0},
		20, 2.5,
	}

	fmt.Println(aRect.Name)
	fmt.Println(aRect.Shape)
	fmt.Println(aRect.Shape.Name)
}

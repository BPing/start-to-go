package _package

import "fmt"

// 1、在同一个package中，可以多个文件中定义init方法
// 2、在同一个go文件中，可以重复定义init方法
// 3、在同一个package中，不同文件中的init方法的执行按照文件名先后执行各个文件中的init方法
// 4、在同一个文件中的多个init方法，按照在代码中编写的顺序依次执行不同的init方法

func init() {
	fmt.Println("init")
}

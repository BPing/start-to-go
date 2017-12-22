package array

import "fmt"

// https://blog.go-zh.org/go-slices-usage-and-internals
//
//
// type slice struct {
//        len int
//        cap int
//        ptr *Elem
//  }
//

func EchoSlice() {
	fmt.Println("EchoSlice")
	arrays := [5]int{1, 2, 3, 4, 5} //这是一个数组，指明长度了
	slices := []int{1, 2, 3}        //这是一个slice
	//slices2 := make([]int, 3)       //这样也可以是一个slice
	slices3 := arrays[1:3] // 是一个slice slice[low:high] 包括low但不包括high，即是 low<=x<high
	slices4 := arrays[1:3] // 是一个slice slice[low:high] 包括low但不包括high，即是 low<=x<high
	slices3[0] = 12
	fmt.Printf("slices3=>%v \n", slices3)
	fmt.Printf("slices4=>%v \n", slices4)

	modifyArr(arrays)
	fmt.Printf("arrays=>%v \n", arrays)

	modifySlice(slices)
	fmt.Printf("slices=>%v \n", slices)
}

// 值复制
func modifyArr(arr [5]int) {
	arr[0] = 10
}

// 引用
func modifySlice(slice []int) {
	slice[0] = 10
}

func ExtendSlice() {
	fmt.Println("ExtendSlice")
	slices := []int{1, 2, 3} //这是一个slice

	fmt.Printf("cap=> %d len=> %d \n", cap(slices), len(slices))

	slices = append(slices, 1)
	fmt.Printf("cap=> %d len=> %d \n", cap(slices), len(slices))

}

func CopySlice() {
	fmt.Println("CopySlice")
	slices := []int{1, 2, 3}  //这是一个slice
	slices2 := make([]int, 2) //这样也可以是一个slice

	copy(slices, slices2)
	fmt.Printf("cap=> %d len=> %d \n val=>%v \n", cap(slices), len(slices), slices)
	fmt.Printf("cap=> %d len=> %d \n val=>%v \n", cap(slices2), len(slices2), slices2)
}

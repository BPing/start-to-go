package array

import "fmt"

func EchoArray() {
	// 定义初始数组
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrAgeZore = [5]int{18, 20, 15}
	var arrLazy = [...]int{5, 6, 7, 8, 22}

	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("[arrLazy] Person at %d is %s\n", i, arrAge[i])
	}

	for k, val := range arrAgeZore {
		fmt.Printf("[arrAgeZore] Person at %d is %s\n", k, val)
	}

	for i := 0; i < len(arrLazy); i++ {
		fmt.Printf("[arrLazy] Person at %d is %s\n", i, arrLazy[i])
	}
}

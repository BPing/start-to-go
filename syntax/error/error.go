package error

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
		//return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func EchoError() {
	_, err := Sqrt(-1)
	if err != nil {
		// 调用 (Error() string) 方法输出信息
		fmt.Println(err)
	}
}

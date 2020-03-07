package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func enval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// div 带余， 除法
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// apply 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)", opName, a, b)
	return op(a, b)
}

// sum 使用可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// main 主函数
func main() {
	fmt.Println(enval(1, 3, "+"))
	fmt.Println("Error handling:")
	if result, err := enval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(10, 3)
	fmt.Printf("10 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is :", apply(
		func(a, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5=", sum(1, 2, 3, 4, 5))
}

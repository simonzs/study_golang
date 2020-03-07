package main

import (
	// "net/http"

	"fmt"
	"math"
	"math/cmplx"
)

// Variable 变量
func Variable() {
	var i int
	var s string
	fmt.Println(i, s)
	fmt.Printf("%d %q\n", i, s)
}

// VariableInitialValue 初始化数值
func VariableInitialValue() {
	var a, b int = 3, 4
	fmt.Println(a, b)
}

// VariableTypeDeduction 推断
func VariableTypeDeduction() {
	var a, b = 1, true
	fmt.Println(a, b)
}

// VariableShorter 简单写法
func VariableShorter() {
	a, b := 1, 3
	fmt.Println(a, b)
}

// euler 欧拉公式 e^(i*PI) + 1 = 0
func euler() {
	var c complex128
	c = 3 + 4i
	fmt.Println(c)
	fmt.Println(cmplx.Abs(c))
	fmt.Println("e^(i*PI) + 1 = ", cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Println("e^(i*PI) + 1 = ", cmplx.Exp(1i*math.Pi)+1)
}

// triangle 不能强制类型转换
func triangle(a, b int) int {
	return calcTriangle(a, b)
}

// calcTriangle 转换
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//consts 常量的定义
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = calcTriangle(a, b)
	fmt.Println(c)
}

// enums 枚举类型
func enums1() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

// main 主函数
func main() {
	fmt.Println("hello world")
	Variable()
	VariableInitialValue()
	VariableShorter()
	euler()
	triangle(3, 4)
	consts()
	enums1()
}

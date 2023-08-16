package main

import (
	"errors"
	"fmt"
)

func intSum(x int, y int) int {
	return x + y
}

func intSum1(x, y int) int {
	return x + y
}

// 可变参数
func intsum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

func calc3() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}

// 函数作为参数
func calc4(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别")
		return nil, err

	}
}

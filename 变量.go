package main

import "fmt"

func main1() {
	// 声明一个遍历
	var name0 string
	fmt.Println(name0)

	// 声明多个遍历
	var (
		name string
		age  int
	)

	name = "pcw"
	age = 20
	fmt.Println(name, age)

	// 声明变量直接赋值
	var addr string = "gd"
	fmt.Println(addr)

	// 声明多个遍历并且赋值
	var name1, age1 = "q1mi", 20
	fmt.Println(name1, age1)

	var name3 = "Q1mi"
	var age3 = 18
	fmt.Println(name3, age3)

	fmt.Println("hello world!")

}

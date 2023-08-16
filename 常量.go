package main

import "fmt"

func main2() {
	const pi0 = 3.1415
	const e0 = 2.7182

	fmt.Println(pi0, e0)

	const (
		pi = 3.1415
		e  = 2.7182
	)

	fmt.Println(pi, e)
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	//iota  记录常量
	const (
		n11 = iota * 10 //0
		n12             //1
		n13             //2
		n4              //3
	)
	fmt.Println(n11, n12, n13, n4)

}

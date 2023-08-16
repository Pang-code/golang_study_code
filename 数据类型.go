package main

import (
	"fmt"
	"math"
)

func teshu() {
	v := 0b00101101 //  是二进制的101101  是十进制的45
	v1 := 0o377     // 八进制的377  十进制的255
	v3 := 0x1p-2    //十六进制   十进制的0.25
	v4 := 123_456

	fmt.Println(v, v1, v3, v4)

}

func intType() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
}

func floatType() {
	fmt.Printf("%f\n", math.Pi)   //
	fmt.Printf("%.2f\n", math.Pi) // 保留两位小数
}

func complexType() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}

//func boolType() {
//	t = true
//	f = false
//	if t == f {
//		fmt.Println("true")
//	}
//
//}

func strType() {
	s1 := "hello"
	s2 := "你好"

	s3 := `slskj
		lskjlsk
		lskdjflkjsdfkj
		lkj`

	fmt.Println(s1, s2, s3)

	fmt.Println("s1的长度为：", len(s1))

	//拼接
	fmt.Println("s1的长度为：" + "3333333333333")

	//len(str)	求长度
	//+或fmt.Sprintf	拼接字符串
	//strings.Split	分割
	//strings.contains	判断是否包含
	//strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	//strings.Index(),strings.LastIndex()	子串出现的位置
	//strings.Join(a[]string, sep string)	join操作

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

package main

import (
	"fmt"
)

//map[KeyType]ValueType

func mapDemo() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo) //

	//	判断key是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

func mapDemo1() {
	scormap := make(map[string]int)
	scormap["张三"] = 90
	scormap["小明"] = 100
	scormap["娜扎"] = 60
	for k, v := range scormap {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range scormap {
		fmt.Println(k)
	}

	delete(scormap, "小明")

	fmt.Println(scormap)
}

package main

import "fmt"

func main() {

	//调用变量
	main1()

	//调用常量
	main2()

	teshu()

	intType()

	floatType()

	complexType()

	//boolType()

	strType()

	traversalString()

	ifDemo1()

	ifDemo2()

	forDemo()

	forDemo2()

	forDemo3()

	switchDemo1()

	testSwitch3()

	switchdemo4(20)

	switchDemo5()

	gotodemo1()

	gotoDemo2()

	breakDemo1()

	continueDemo()

	arraydemo()

	arraydemo1()

	arraydemo2()

	arraydemo3()

	arraydemo4()

	arraydemo5()

	arraydemo6()

	a := [3]int{1, 2, 3}
	s := arraySum(a)
	fmt.Println(s)

	appenddemo()

	mapDemo()

	mapDemo1()

	sum := intSum(1, 3)
	fmt.Println(sum)

	ret4 := intsum2(10, 20, 30)
	fmt.Println(ret4)

	calc1(1, 2)
	calc2(1, 2)
	calc3()

	ret5 := calc4(1, 5, add)
	fmt.Println(ret5)

	f, e := do("+")
	if e == nil {
		ret6 := f(9, 9)
		fmt.Println(ret6)
	} else {
		fmt.Println("输入错误")
	}

}

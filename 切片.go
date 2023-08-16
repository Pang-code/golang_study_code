package main

func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//func main() {
//	a := [5]int{1, 2, 3, 4, 5}
//	s := a[1:3] // s := a[low:high]
//	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
//}

//func main() {
//	a := make([]int, 2, 10)
//	fmt.Println(a)      //[0 0]
//	fmt.Println(len(a)) //2
//	fmt.Println(cap(a)) //10
//}

func appenddemo() {
	var s []int
	s = append(s, 1)       // [1]
	s = append(s, 2, 3, 4) // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]
}

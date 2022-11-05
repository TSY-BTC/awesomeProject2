package main

import "fmt"

func main() {

	a := [2]int{1, 2}
	fmt.Printf("%T\n", a)

	// 指针类型
	b := 100
	p := &b
	fmt.Printf("%T\n", p)

	// 切片类型
	c := []int{1, 23, 4, 3, 67, 3}
	fmt.Printf("%T\n", c)

	test()

	fmt.Printf("funck")

	fmt.Printf("bug2 改好了")

	fmt.Printf("修改1")

	fmt.Printf("GITHUB 推送       ")

	fmt.Printf("10.12 test")

	/* 	打印类型
	   	var name string = "tome"
	   	age := 20
	   	b := true
	   	fmt.Printf("%T\n", name)
	   	fmt.Printf("%T\n", age)
	   	fmt.Printf("%T\n", b)
	*/

}

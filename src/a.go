package main

import "fmt"

func test() {
	/*	//常量的使用
		const PI float32 = 3.14
		const (
			a = 100
			b = 200
		)
		print(a)*/

	// iota 的使用
	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("%v\n%v\n%v", a1, a2, a3)
}

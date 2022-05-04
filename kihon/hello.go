package main

import (
	"fmt"
)

func main() {
	var (
		aaa   string = "aaa"
		hello        = "hello"
	)
	fmt.Println(aaa)
	fmt.Println(hello)

	ifFunc(10, 100)
	forFunc(3)
	switchFunc(5)
}

// if分岐の書き方
func ifFunc(a int, b int) {
	if a > b {
		fmt.Println("aのほうがでかい")
	} else if b > a {
		fmt.Println("bのほうがでかい")
	} else {
		fmt.Println("同じ値")
	}
}

// for文の書き方
func forFunc(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
	fmt.Printf("%d 回目のループはbreak\n", n)
}

// switch文の書き方
func switchFunc(n int) {
	switch n {
	case 10:
		fmt.Println("10です")
	case 5, 9:
		fmt.Println("5 or 9 です")
	default:
		fmt.Println("それ以外です")
	}
}

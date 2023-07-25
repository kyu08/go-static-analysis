package a

import "fmt"

func F() {
	hoge := "str"
	fmt.Printf("hoge: %v\n", hoge)
	res := 1 + 2
	print(res)
}

var a = 1 // want "a should be more than 3 characters"

const hoge = 2

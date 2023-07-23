package a

import "fmt"

func F() {
	hoge := "str"
	fmt.Printf("hoge: %v\n", hoge)
	res := 1 + 2 // want "int found" "int found"
	print(res)   // want "int found"
}

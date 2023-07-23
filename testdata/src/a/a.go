package a

func F() {
	var gopher int
	print(gopher)
}

func unusedFunc() {} // want `"unusedFunc" is unused`
func PublicFunc() {}

func main() {}
func init() {}

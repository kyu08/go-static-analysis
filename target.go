package main

type user struct{}

func (user) M1() {}
func (user) M2() {}

func Hoge(name string) (user, error) {
	u := user{}
	u.M1()
	privateFunc()

	return u, nil
}

func Func1() bool   { return true }
func Func2() int    { return 2 }
func Func3() string { return "a" }
func Func4() string { return "" }

func privateFunc() string { return "" }

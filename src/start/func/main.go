package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) (result, multiple int) {
	multiple = a * b
	result = a + b
	return
}

func sum(a ...int) (ret int) {
	fmt.Printf("%T\n", a)
	for _, v := range a {
		ret += v
	}
	return
}

func sum01(b string, a ...int) (ret int) {
	fmt.Printf("%T\n", a)
	for _, v := range a {
		ret += v
	}
	return
}

func add00(a int, b int) (result int) {
	result = a + b
	return
}

func add01(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func cloo(name string) func() {
	return func() {
		fmt.Println("\nclosure", name)
	}
}

func cloo02(name string) func(string) {
	return func(_name string) {
		fmt.Println("\nclosure", _name)
	}
}

func prefixString(prefix string) func(string) string {
	return func(str string) string {
		if !strings.HasPrefix(str, prefix) {
			return prefix + str
		}
		return str
	}
}

func aass(base int) (func(int) int, func(int) int) {
	a := func(addItem int) int {
		return base + addItem
	}

	s := func(subItem int) int {
		return base - subItem
	}

	return a, s
}

func panicDefer(name string) {
	defer func(name string) {
		fmt.Println("defer", name, recover())
	}(name)
	panic("panicDefer")
}

func panicDeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err", err)
		}
	}()
	panic("panicDefer")
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3, 6, 3, 4))

	fmt.Println(sum01("ee", 6, 3, 4))
	fmt.Println(add(3, 3))

	ff := sum
	fmt.Printf("%T", ff)
	fmt.Println(ff(3, 3, 7))

	fmt.Println(add01(100, 33, add00))
	fmt.Println("ee")

	sayFunc := func() {
		fmt.Println("anonymouse")
	}

	sayFunc()

	func(a int, b string) {
		fmt.Printf("%02d %s", a, b)
	}(3, "eee")

	cc := cloo("eeesssss")
	cc()

	pre := prefixString("prefixxx")
	fmt.Println(pre("eeee"))
	fmt.Println(pre("prefixxxeeee"))
	panicDeferRecover()
	aa, ss := aass(100)
	fmt.Println(aa(50), ss(50))

}

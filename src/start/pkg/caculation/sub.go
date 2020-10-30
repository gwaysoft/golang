package caculation

import "fmt"

func Sub(x, y int) int {
	fmt.Println(Add(x, y))
	return sub(x, y)
}

func sub(x, y int) int {
	return x - y - y
}

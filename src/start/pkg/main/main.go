package main

import (
	"fmt"
	"start/pkg/caculation"
	c "start/pkg/caculation"
)

func main() {
	fmt.Println(caculation.Add(7, 8))
	fmt.Println(caculation.Sub(7, 8))
	fmt.Println(c.Add(7, 8))
}

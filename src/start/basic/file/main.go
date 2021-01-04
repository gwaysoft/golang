package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./t.tt")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer file.Close()

	var tmp = make([]byte, 10, 10)

	n, err := file.Read(tmp)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("%d", n)
	fmt.Println(string(tmp))
}

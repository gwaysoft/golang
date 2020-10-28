package main

import(
	"fmt"
)

type Person struct {
	name, city string
	age int
}

func main()  {
	fmt.Println("eee")
	p := Person{name: "da", city: "shanghai",age:24}
	fmt.Println(p)
}

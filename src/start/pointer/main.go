package main

import (
	"fmt"
)

func ppp(p *int) int {
	*p *= 2
	return *p
}

type myInt int

func main() {
	// fmt.Println("g")
	a := 5
	b := &a

	fmt.Printf("%003d, %p\n", a, b)
	fmt.Printf("%T, %T\n", a, b)
	fmt.Printf("%v, %v\n", b, &a)
	fmt.Printf("%v\n", *b)
	fmt.Printf("%v, %p\n", b, b)
	fmt.Printf("%v, %p\n", b, &b)

	var c *int = &a
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", *c)
	*c = 100
	fmt.Println(a)

	p := 50

	fmt.Println(ppp(&p))
	fmt.Println(p)

	var ap *int
	ap = new(int)
	// ap := new(int)
	fmt.Printf("%T %v\n", ap, *ap)
	*ap = 1000
	fmt.Printf("%T %v\n", ap, *ap)

	var mp map[string]int
	mp = make(map[string]int, 8)
	mp["点点滴滴"] = 5
	mp["点点滴滴 "] = 5
	fmt.Printf("%#v\n", mp)

	var myT myInt // = 5
	myT = 5
	fmt.Printf("%T %v %p\n", myT, myT, &myT)

	type aliasType int
	var myAT aliasType
	myAT = 56
	fmt.Printf("%T %v %p\n", myAT, myAT, &myAT)

	fmt.Printf("%T %v %p\n", a, a, &a)

}

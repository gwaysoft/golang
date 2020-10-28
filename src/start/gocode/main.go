package main

import "fmt"
import "strings"
import "sort"

// import "make"

var global_name = "global_name"

func main() {
	var (
		a    int    = 18
		name string = "dddddd噩噩噩噩"
	)
	var name2, age = "eee", 32
	m := "666"
	fmt.Println(a, name, name2, age, m)
	fmt.Println(global_name)

	const (
		ac = iota
		b  = 100
		c  = iota
	)
	fmt.Println(ac, b, c)

	fmt.Println(1 << 10)

	if i := 0; i > 10 {
		fmt.Println("a")
	} else if i < 10 && i > 0 {
		fmt.Println("b")
	} else if i < 0 {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}

	for i := 0; i < 10; i++ {

		fmt.Printf("%b\n", i)
		// switch {
		// case (i % 2) == 0:
		// 	fmt.Println(i, "o")
		// case (i % 2) == 1:
		// 	fmt.Println(i, "J")
		// default:
		// 	fmt.Println("dd")
		// }

		// if i == 8 {
		// 	break
		// }

		// if i < 5 {
		// 	continue
		// }
		// fmt.Println(i)

	}

	var bb bool
	fmt.Println(bb)

	var s3 string = `
	fei
	dde
	\nsss
	ee
	    woww 
	`
	fmt.Println(strings.Split(s3, "\n"))

	fmt.Printf("%T", strings.Split(s3, "\n"))

	var rrr rune = '我'
	fmt.Println(rrr)
	fmt.Printf("%T \n", rrr)

	var s2 string = "wwdes吊顶"

	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c\n", s2[i])
	}

	for _, r := range s2 {
		fmt.Printf("%c\n", r)
	}

	var aa = [3]string{"3", "3x32"}
	fmt.Println(aa)

	var ii = [...]int{3, 2, 3, 5, 2}
	fmt.Println(ii, len(ii))

	// var acc [3]int;

	var strr = [...]string{3: "3333", 9: "333"}
	fmt.Println(strr)
	fmt.Printf("%T\n", strr)

	for in, va := range strr {
		fmt.Println(in, va)
	}

	var slice []int
	if slice == nil {
		fmt.Println("slice == nil")
	}

	fmt.Println(slice)

	// iii := ii[:len(ii)-1]
	iii := []int{3, 4, 10000}
	slice = append(slice, 2, 3, 4, 5)
	slice = append(slice, iii...)

	fmt.Println(slice)

	slice01 := []int{}

	slice01 = []int{2}
	fmt.Println(slice01)
	if slice01 == nil {
		fmt.Println("slice01 == nil")
	}

	var bs = []bool{false, true}
	fmt.Println(bs)

	ar := [5]int{2, 3, 4, 2, 3}
	// as := ar[1:4]
	// acc := as[0:len(as)]
	acc := ar[:len(ar)-1]
	fmt.Println(acc)
	fmt.Println(cap(acc))

	for index, value := range acc {
		fmt.Println(index, value)
	}

	var ccs = make([]int, 3, 5)
	ccs[2] = 100
	fmt.Println(ccs)

	var a01 []int

	for i := 1; i < 10; i++ {
		a01 = append(a01, i)
		fmt.Printf("%p %d %d %v \n", a01, len(a01), cap(a01), a01)
	}

	var a02 = make([]int, 9, 9)
	copy(a02, a01)

	fmt.Println(a01)
	a01[0] = 10
	fmt.Println(a01)
	fmt.Println(a02)

	index01 := 3

	a02 = append(a02[:index01], a02[index01+1:]...)
	fmt.Println(a02)

	a03 := [...]int{3, 5, 2, 5, 30, 4}
	sort.Ints(a03[:])
	fmt.Println(a03)
}

package main

import (
	"fmt"
)

type person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) person {
	return person{
		name: name,
		city: city,
		age:  age,
	}
}

func newPersonP(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func (p person) kaka(a int) int {
	p.name = "kaka"
	fmt.Println("kaka: p.name = ", p.name)
	return a * a
}

func (p *person) kakaPP(name string) string {
	p.name = name
	fmt.Println("kakaPP: p.name = ", p.name)
	return name
}

func main() {
	p := person{name: "33", city: "ee", age: 10}
	// var p person
	// p.name = "ee"
	// p.city = "d吊顶"
	// p.age = 14
	fmt.Printf("%T %#v %p\n", p, p, &p)

	// anonymousS := struct{name: "ee", age:26}
	var anonymousS struct {
		name string
		age  int
	}
	anonymousS.name = "eeee"
	anonymousS.age = 24
	fmt.Printf("%T %#v %p\n", anonymousS, anonymousS, &anonymousS)

	// p2 := new(person)
	var p2 = new(person)
	(*p2).name = "ee"
	p2.city = "ckg"
	p2.age = 24
	fmt.Printf("%T %#v\n", p2, p2)

	// var p3 person
	var p3 = person{}
	// p3 := &person{}
	// p3 := new(person)
	// p3 := person{}
	p3.name = "ee"
	p3.city = "d吊顶"
	p3.age = 14
	fmt.Printf("%T %#v %p\n", p3, p3, &p3)
	fmt.Println(p3.kaka(5000))

	cp := newPerson("ee", "ee3", 12)
	cpC := cp
	cp.name = "newee"

	fmt.Printf("%T %#v %p\n", cp, cp, &cp)
	fmt.Printf("%T %#v %p\n", cpC, cpC, &cpC)

	cpP := newPersonP("ee", "ee3", 12)
	cpP01 := cpP
	cpP.name = "newee"

	fmt.Printf("%T %#v %p\n", cpP, cpP, cpP)
	fmt.Printf("%T %#v %p\n", cpP01, cpP01, cpP01)

	fmt.Println(cpP.kaka(2))
	fmt.Printf("%T %#v %p\n", cpP, cpP, cpP)

	fmt.Println(cpP.kakaPP("kakaPPP"))
	fmt.Printf("%T %#v %p\n", cpP, cpP, cpP)

}

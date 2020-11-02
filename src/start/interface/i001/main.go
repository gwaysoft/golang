package main

import "fmt"

type sayer interface {
	say(string) string
}

type ant struct{}

type bee struct{}

func (a ant) say(str string) string {
	fmt.Println(str)
	return str + "return"
}

func (b bee) say(str string) string {
	fmt.Println(str)
	return str + "return"
}

func say(insect sayer, content string) string {
	return insect.say(content)
}

type pig struct {
	name string
}

func (p pig) say(string2 string) string {
	fmt.Println(p.name, string2)
	return p.name + string2
}

func main() {
	println(say(new(ant), "ant new"))
	println(say(new(bee), "bee new"))

	fmt.Println(say(pig{name: "pigggg"}, "pink"))

	println("eeeeeeeeeeeeee")
	println(say(ant{}, "ant{}"))
	println(say(bee{}, "bee{}"))

	var sss sayer
	sss = new(ant)
	fmt.Println(sss.say("eee"))

}

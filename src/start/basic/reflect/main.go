package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string `ini:"name" json:"jname"`
	age  int8   `ini:"age" json:"jage"`
}

// capital -> public
func (c cat) Jiao() {
	println(c.Name)
}

func reflectType(xx interface{}) {
	var t = reflect.TypeOf(xx)
	fmt.Println(t, t.Name(), t.Kind())
}
func main() {
	reflectType(5)
	reflectType([]int{1, 3})
	reflectType(cat{Name: "ee"})
	reflectType(cat{Name: "ee"}.Jiao)
	c := cat{Name: "ee", age: 12}
	t := reflect.TypeOf(c)

	fmt.Println(t.NumMethod())
	fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(i, f.Name, f.Tag, f.Type)
		fmt.Println(f.Tag.Get("ini"), f.Tag.Get("json"))
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(i, m.Name, m.Type, &m.Func)
	}
}

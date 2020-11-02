package main

import "fmt"

type roller interface {
	roll(string)
}

type valuer interface {
	value(float32)
}

type base interface {
	roller
	valuer
}

type stone struct {
	name string
	qian float32
}

func (s stone) roll(localName string) {
	fmt.Println("roll -> ", s.name, localName, s.qian)
}

func (s stone) value(f float32) {
	fmt.Println("value -> ", f, s.qian)
}

//func (s *stone) value(f float32) {
//	fmt.Println("value -> ", f, s.qian)
//}

//func (s stone)base()  {
//	roller
//	valuer
//}

//func (s *stone)roll(localName string)  {
//	fmt.Println("roll -> ", s.name, localName)
//}

type null interface{}

func nullFunc(nu ...interface{}) {
	fmt.Printf("%v, %T\n", nu, nu)
}

func main() {
	s := stone{name: "001-stone"}
	//s := &stone{name: "pointer-stone"}
	var r roller
	r = s
	//r = &s
	s.qian = 90.22
	s.name = "002"
	r.roll("localllll")

	var v valuer
	v = &s
	s.qian = 0.901
	v.value(500)

	var b base
	//s = stone{qian: 55.3, name: "basccccc"}
	//b = s
	b = &s
	//s = stone{qian: 55.3, name: "basccccc"}
	b.value(333.33)
	b.roll("basic")

	var n null = struct {
		name string
	}{"ee"}
	fmt.Println(n)

	nullFunc("ddd", 12)

	//var m  = make(map[string]interface{}, 12)
	m := make(map[string]interface{}, 12)
	m["name"] = "dd"
	m["age"] = 12
	fmt.Printf("%v\n", m)

	//vm := map[string]interface{}{"name" :12}
	var vm = map[string]interface{}{"name": 12}
	vm2 := vm
	vm["name"] = "ddwww"
	fmt.Printf("%v\n", vm2)

	ret, ok := vm["name"].(bool)
	if !ok {
		fmt.Println("ret", ret)
	}

	switch t := vm["name"].(type) {
	case bool:
		fmt.Println("bool", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("default", t)
	}

}

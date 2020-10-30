package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	province string
	City     string
}

func (a *Address) location(province string) {
	a.province = province
	a.City = "change"
	println("a l", a.province, a.City)
}

type email struct {
	id  int
	Mmm string `json:"address"`
}

func newEmail(id int, mmm string) *email {
	return &email{id: id, Mmm: mmm}
}

func newEmail1(id int, mmm string) email {
	return email{id: id, Mmm: mmm}
}

type Person struct {
	// for json
	Name    string  `json:"nameee"`
	address Address
	*Address
	Email []email `json:"emails" db:"dbemails" xml:"xmlemails"`
	age   int
	story struct {
		s int
	}
}

func (p *Person) mm() {
	println("mm = ", p.Name)
}

func main() {

	p := Person{Name: "person_name", address: Address{province: "DS", City: "Qd"},
		age: 23, story: struct{ s int }{s: 12}, Address: &Address{province: "ee", City: "eee"}}

	//p.Email[0] = email{Mmm: "eee"}
	//p.Email = []email{{Mmm: "string(ee)"}, {Mmm: "eeeee01"}, *newEmail(1, "ddd"), newEmail1(3, "33333")}

	p.Email = make([]email, 5, 5)
	for i := 0; i < len(p.Email); i++ {

		p.Email[i] = email{id: i, Mmm: fmt.Sprintf("sss%02d", i)}
		//k := fmt.Sprintf("stu%02d", i)
		//p.Email[i].Mmm = k
	}
	fmt.Printf("%#v %d %T\n", p.Email, len(p.Email), p.Email)
	// Capital is public, and lower is viewed at the same package
	data, err := json.Marshal(p)
	if err == nil {
		fmt.Printf("%#v\n", p)
		fmt.Printf("%#v %T\n", p.Address, p.Address)
		fmt.Printf("%s\n", data)
	}else {
		println(err)
	}

	//str := `{"Name":"kakaperson_name","age":13,"Email":[{"Mmm":"sss00"},{"Mmm":"sss01"},{"Mmm":"sss02"},{"Mmm":"sss03"},{"Mmm":"sss04"}]}`
	str := `{"nameee":"strperson_name","City":"1111eee","age":13,"Address":{"City":"ss", "province":"SSSS"},"Email":[{"Mmm":"sss00"},{"Mmm":"sss01"},{"Mmm":"sss02"},{"Mmm":"sss03"},{"Mmm":"sss04"}]}`

	var p2 Person
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		println("eeeeeeeeer",err)
	}else {
		fmt.Printf("%#v\n", p2)
		fmt.Printf("%v %v\n", p2, p2.Address)
	}
	println(p.address.province, p.address.City)
	p.address.location("Ws")
	println(p.address.province, p.address.City)

	println(p.province, p.City)
	p.location("kakak")
	println(p.province, p.City)

	p.mm()

}

package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("1. add; 2. edit; 3. show; 4. exit")
}

func inputStudent() *student {
	var (
		id          int
		name, class string
	)
	fmt.Print("Input Id: ")
	i, err := fmt.Scanf("%d\n", &id)
	//fmt.Println(i, err)
	defer func() {
		fmt.Println(i, err)
		if err := recover(); err != nil{
			println("dd",err)
		}
	}()
	println("3333333333", err)
	if err != nil{
		println("aaaaa")
		panic(err)
	}




	fmt.Print("Input name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Input class: ")
	fmt.Scanf("%s\n", &class)
	return newStudent(id, name, class)
}
func main() {
	showMenu()
	studentManager := newStudentManager()
	//var a, b bool
	var a bool
	for {
		var input int
		fmt.Println(a)
		if a {
			a = false
			fmt.Scanf("\n")
			continue
		}
		fmt.Print("Input number: ")

		_, err := fmt.Scanf("%d\n", &input)
		fmt.Println("EX: ", err, input)
		if err != nil {
			a = true
			continue
		}
		//a = false
		fmt.Printf("Input number is %d\n", input)

		switch input {
		case 0:
			continue
		case 1:
			student := inputStudent()
			fmt.Printf("%v\n" +
				"", student)
			if student == nil {
				println("st01: ", student)
				a = true
				continue
			}
			//b = false
			studentManager.addStudent(student)
		case 2:
			student := inputStudent()
			studentManager.editStudent(student)
		case 3:
			studentManager.showStudent()
		case 4:
			os.Exit(0)
		}
	}
}

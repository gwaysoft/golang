package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name string, class string) (ret *student) {
	ret = &student{id: id, name: name, class: class}
	return
}

type studentManager struct {
	allStudents []*student
}

func newStudentManager() *studentManager {
	ret := &studentManager{
		allStudents: []*student{},
	}
	return ret
}

func (s *studentManager) addStudent(student *student) *student {
	s.allStudents = append(s.allStudents, student)
	return student
}

func (s *studentManager) editStudent(student *student) *student {
	for i, v := range s.allStudents {
		if student.id == v.id {
			s.allStudents[i] = student
			return student
		}
	}
	fmt.Println("nothing to edit")
	return nil
}

func (s *studentManager) showStudent() {
	fmt.Println("show student list", len(s.allStudents))
	for _, v := range s.allStudents {
		fmt.Printf("id %#v\n", v)
	}
}

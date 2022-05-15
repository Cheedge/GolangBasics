package main

import "fmt"

type Person struct {
	age int
	sex string
}

type Student struct {
	p     Person
	grade int
}

func callp(p *Person) {
	p.age = 100
	p.sex = "m"
	fmt.Println("in callp func")
	fmt.Printf("%p, age=%d, sex=%s\n", p, p.age, p.sex)
}

func callstu(s *Student) {
	s.p.age = 300
	s.p.sex = "c"
	s.grade = 6
	fmt.Println("in callstu func")
	fmt.Printf("%p, age=%d, sex=%s\n", s, s.p.age, s.p.sex)
}

/*
不能返回局部变量的地址，
因为函数结束，stack frame释放，
但局部变量所在空间不会消失，会重新分配，
所以地址会变，但其值不会改变
*/

func returnpassparams(s *Student) *Student {
	// s = new(Student)
	s.p.age = 10000
	s.p.sex = "noidea"
	s.grade = 999
	fmt.Printf("%p, age=%d, sex=%s\n", s, s.p.age, s.p.sex)
	fmt.Printf("&s INSIDE FUNC IS %p\n", &s)
	return s
}

func main() {
	// p := Person{
	// 	age: 1,
	// 	sex: "f",
	// }
	var p = new(Person) // new() return a ptr
	fmt.Println("before call")
	fmt.Printf("%p, age=%d, sex=%s\n", p, p.age, p.sex)
	callp(p)
	fmt.Println("after call")
	fmt.Printf("%p, age=%d, sex=%s\n", p, p.age, p.sex)

	var s = new(Student)
	callstu(s)
	fmt.Println("after callstudent")
	fmt.Printf("%p, age=%d, sex=%s\n", s, s.p.age, s.p.sex)

	var s1 = new(Student)
	s1 = returnpassparams(s1)
	fmt.Println("pass params by return")
	fmt.Printf("%p, age=%d, sex=%s\n", s1, s1.p.age, s1.p.sex)
	fmt.Printf("&s CHANGED TO %p\n", &s)
}

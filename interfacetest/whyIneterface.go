package main

import (
	"fmt"
)

type Cat struct{}
type Dog struct{}

func (c Cat) speak() {
	fmt.Println("Meow, Meow")
}

func (d Dog) speak() {
	fmt.Println("Wang, Wang")
}

type Speaker interface {
	speak()
}

func talk(s Speaker) {
	s.speak()
}

func main() {
	c := Cat{}
	d := Dog{}
	// c.meow()
	// d.bark()
	talk(c)
	talk(d)

	var mao Cat
	gou := Dog{}
	var spk Speaker
	spk = mao
	spk = gou
	fmt.Println(spk)

}

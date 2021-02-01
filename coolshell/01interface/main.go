package main

import "fmt"

type PersonInf interface {
	Age1(age int) int
	Age2(age int) int
}

type Person struct {
	age int
}

type A struct{}

func (a *A) Age1(age int) int {
	return age
}
func (a *A) Age2(age int) int {
	return age
}

func (s *Person) Age1(age int) int {
	return age
}
func (s *Person) Age2(age int) int {
	return age
}

func main() {
	p := &Person{age: 5}
	fmt.Printf("age:%d", p.Age1(p.age))
}

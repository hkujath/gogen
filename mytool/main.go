package main

import "fmt"

//go:generate gogen ..\datatypes\stack.gogen MyType
//go:generate gogen ..\datatypes\stack.gogen int
func main() {
	mts := MyTypeStack{}
	mts.push(MyType{"foo1", "bar1"})
	mts.push(MyType{"foo2", "bar2"})
	p := mts.Pop()
	fmt.Printf("%#v\n", p)
	is := intStack{}
	is.push(1)
	is.push(2)
	is.push(3)
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
}

type MyType struct {
	Foo string
	Bar string
}

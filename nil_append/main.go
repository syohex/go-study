package main

import "fmt"

type Foo struct {
	a []*int
}

func (f *Foo) funcB(a ...*int) {
	f.a = append(f.a, a...)
}

func funcA(a *int) *Foo {
	f := new(Foo)
	f.funcB(a)
	return f
}

func main() {
	f := funcA(nil)
	fmt.Printf("%+v\n", f)
	fmt.Printf("len=%d\n", len(f.a))
}

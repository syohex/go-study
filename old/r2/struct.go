package main

import (
	"fmt"
)

type Sample struct {
	Name  string
	Age   int
}

func (s *Sample) Hello() {
	fmt.Printf("Hello %s, age %d\n", s.Name, s.Age)
}

func main() {
	tms := Sample{"TMS", 20}
	tms.Hello()
}

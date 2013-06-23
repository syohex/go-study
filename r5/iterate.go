package main

import "fmt"
import "time"

func main () {
	var a [100]int

	for i := 1; i < 10; i++ {
		fmt.Printf("Element %d is %d\n", i, a[i])
	}

	subrange := a[1:10]
	for i, v := range subrange {
		fmt.Printf("Element: %d %d\n", i, v)
	}

	for i, v := range subrange {
		go fmt.Printf("Element: %d %d\n", i, v)
	}
	time.Sleep(1000000000)
}

package main

import "fmt"

func main() {

	fmt.Println("for with 1 condition")
	i := 1
	for i <= 3 {
			fmt.Println(i)
			i = i + 1
	}

	fmt.Println("for with init, condition, iterator")
	for j := 0; j < 3; j++ {
			fmt.Println(j)
	}

	fmt.Println("for with range")
	for i := range 3 {
			fmt.Println("range", i)
	}

	fmt.Println("for only")
	for {
			fmt.Println("loop")
			break
	}

	fmt.Println("for with range and continue")
	for n := range 6 {
			if n%2 == 0 {
					continue
			}
			fmt.Println(n)
	}
}
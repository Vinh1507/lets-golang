package main

import "fmt"

func multiple(x *int) {
	*x = *x * 2
}

func main() {
	x := 2
	multiple(&x)
	fmt.Print(x)
}

package main

import "fmt"

//Package level
var a int = 10

const b = "Welcome"

func main() {
	// block level
	c := 25
	fmt.Printf("The value of a is %v and the type of a is %T \n", a, a)
	fmt.Printf("The value of b is %v and the type of b is %T \n", b, b)
	fmt.Printf("The value of c is %v and the type of c is %T", c, c)
}

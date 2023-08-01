package main

import "fmt"

// Everything in go is pass by value

// func(r reciever) identifier(p parameter(s)) (return (s)) {<Your code here>}
func foo(s string) {
	fmt.Println("I am from foo")
	fmt.Println(s)
}
func main() {
	foo("Deba")
	x := aloha("D")
	fmt.Println(x)
	fmt.Println("------------------")
	sum(1, 2, 3, 45, 6, 67)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr ", s)
}

// variadic Parameters - Can pass any number of parameters

func sum(ii ...int) {
	fmt.Println(ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	fmt.Println("Sum is ", n)
}

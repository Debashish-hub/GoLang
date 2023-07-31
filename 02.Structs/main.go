package main

import "fmt"

type person struct {
	first string
	last  string
	favIc []string
}

func main() {
	p1 := person{
		first: "james",
		last:  "Bond",
		favIc: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		favIc: []string{"Strawberry", "vanilla"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("---------------------------\n")

	for _, v := range p1.favIc {
		fmt.Println(p1.first, "favorite is ", v)
	}

	for _, v := range p2.favIc {
		fmt.Println(p2.first, "favorite is ", v)
	}
}

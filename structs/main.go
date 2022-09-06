package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Zeph", lastName: "Gv"}
	fmt.Println(alex)
}

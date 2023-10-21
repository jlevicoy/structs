package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jc := person{"Juan Carlos", "Levicoy Bustamante"}
	fmt.Print(jc.firstName, jc.lastName)
}

package main

import "fmt"

// this syntax is valid. Global varibles are not required to be used.
// this is declarative statement. You are declaring a variable
var globalPiValue = 3.14

// this is not allowed because non-declarative statement is not allowed
//piValue := 3.14

// this is allowed because this is declarative statement
var globaltwoPiValue float64

// this statement is also not allowed becuase statement is non-declarative
//twoPiValue = 2 * 3.14

type linkedListNode struct {
	data int
	next *linkedListNode
}

func main() {
	var age int
	age = 24
	name := "Himanshu" // this can not reused becuase variable has already been created
	fmt.Printf("%s - %d\n", name, age)

	var c, python, goLang = true, false, "yes!" // multiple declaration in single line
	var i, j int = 1, 2                         // declaring type is optional
	fmt.Println(i, j, c, python, goLang)

	// zero value
	var firstName string
	fmt.Println(firstName)
	var salary bool
	fmt.Println(salary)

	var node = linkedListNode{data: 34}
	fmt.Println(node.next == nil)
}

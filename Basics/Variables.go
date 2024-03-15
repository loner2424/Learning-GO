package main

import "fmt"

func variables() {

	var number1 int
	var number2 int = 3
	number3 := 4
	var StringExample = "This is Sahil Sobhani"
	var BoolExample = true

	fmt.Printf("Value of number1 is not initialized, therefore it is %d\n", number1)
	fmt.Printf("Value of number2 is initialized and the type is mentioned, it is %d and an integer\n", number2)
	fmt.Printf("Value of number3 is initialized in a shorter manner and the type is inferred, therefore it is %d\n", number3)
	fmt.Printf("StringExample is a string example, type is inferred, value is %s\n", StringExample)
	fmt.Printf("StringExample is a bool example, type is inferred, value is %t\n\n", BoolExample)

}

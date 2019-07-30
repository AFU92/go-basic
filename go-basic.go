package main

import "fmt"

func main() {

	var num1 float32 = 10
	var num2 float32 = 6
	var operator string = "/"

	helloWorld()

	fmt.Println("Num1 y num2")
	fmt.Println(num1, num2)

	fmt.Print("\nThe operation is " + operator + ":\n")

	fmt.Println(operation(num1, num2, operator))

}

func helloWorld() {
	fmt.Println("Hello World!")
}

func operation(num1 float32, num2 float32, operator string) float32 {
	var result float32
	if operator == "+" {
		result = num1 + num2
	}
	if operator == "-" {
		result = num1 - num2
	}
	if operator == "*" {
		result = num1 * num2
	}
	if operator == "/" {
		result = num1 / num2
	}

	return result
}

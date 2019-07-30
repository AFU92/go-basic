package main

import "fmt"

func main() {

	var num1 float32 = 10
	var num2 float32 = 6

	helloWorld()

	fmt.Print("\nThe sum is: ")

	fmt.Println(operation(num1, num2, "+"))

}

func helloWorld() {
	fmt.Println("Hello World!")
}

func operation(num1 float32, num2 float32, operator string) float32 {
	return num1 + num2
}

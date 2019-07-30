package main

import "fmt"

func main() {

	var num1 float32 = 10
	var num2 float32 = 6

	var num3 int = 10
	var num4 int = 6

	helloWorld()

	fmt.Println("\nNum1")
	fmt.Println(num1)

	fmt.Println("\nNum2")
	fmt.Println(num2)

	calculator(num1, num2)
	calculator2(num3, num4)

	fmt.Println(returnText())
	fmt.Println(returnText2())

}

// Return 1
func returnText() (string, string, int) {
	data1 := "Default"
	data2 := "text"
	data3 := 100
	return data1, data2, data3
}

// Return 2
func returnText2() (data4 string, data5 int) {
	data4 = "Default"
	data5 = 200
	return
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

func operation2(num3 int, num4 int, operator string) int {
	var result int

	if operator == "/" {
		result = num3 / num4
	}
	if operator == "%" {
		result = num3 % num4
	}

	return result
}

func calculator(num1 float32, num2 float32) {
	fmt.Print("\nThe sum is: ")
	fmt.Println(operation(num1, num2, "+"))

	fmt.Print("The rest is: ")
	fmt.Println(operation(num1, num2, "-"))

	fmt.Print("The mult is: ")
	fmt.Println(operation(num1, num2, "*"))

	fmt.Print("The div float is: ")
	fmt.Println(operation(num1, num2, "/"))

}

func calculator2(num3 int, num4 int) {

	fmt.Print("The div int is: ")
	fmt.Println(operation2(num3, num4, "/"))

	fmt.Print("The residue is: ")
	fmt.Println(operation2(num3, num4, "%"))
}

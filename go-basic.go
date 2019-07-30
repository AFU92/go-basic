package main

import "fmt"

func main() {

	fmt.Println(returnText())

	fmt.Println(returnText2())

	fmt.Println(tshirt(45))
}

// recibe una variable de tipo int y devuelve un string y un int
func tshirt(order int) (string, int) {
	return "Quantity t-shit is: ", order
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

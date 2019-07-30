package main

import "fmt"

func main() {

	// fmt.Println(returnText())

	// fmt.Println(tshirt(24, "€"))

	// fmt.Println(returnText2())

	// fmt.Println(tshirt(45, "$"))

	pants("Red", "large", "adidas")

}

func pants(atributes ...string) {

	for _, atribute := range atributes {
		fmt.Println(atribute)
	}

}

// recibe una variable de tipo float32 y devuelve un string y un float32
func tshirt(order float32, money string) (string, string, float32) {

	price := func() float32 {
		return order * 7
	}

	return "The total price is: ", money, price()
}

// Return 1
func returnText() (string, string, int) {
	data1 := "Order"
	data2 := "No."
	data3 := 1
	return data1, data2, data3
}

// Return 2
func returnText2() (data4 string, data5 int) {
	data4 = "Order No."
	data5 = 2
	return
}

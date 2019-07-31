package main

import "fmt"

func main() {
	fmt.Println("Program")

	edad := 25

	if edad >= 18 && edad <= 99 && edad != 25 {
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("Estas en la mejor etapa de tu vida")
	} else if edad > 99 {
		fmt.Println("Eres muy viejo")
	} else {
		fmt.Println("Eres menor de edad")
	}
}

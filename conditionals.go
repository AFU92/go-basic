package main

import (
	"fmt"
	"os"
	"strconv"
)

// correr en la consola y enviar sus paramentros
// go run conditionals.go Andrea 15

func main() {
	fmt.Println("Program")

	fmt.Println("Hola " + os.Args[1] + "\nBienvenido a tu programa")

	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad <= 99 && edad != 25 {
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("Estas en la mejor etapa de tu vida")
	} else if edad > 99 {
		fmt.Println("Eres muy viejo")
	} else {
		fmt.Println("Eres menor de edad")
	}

	numero := 12

	if numero%2 == 0 {
		fmt.Println("Tu edad es par")
	} else {
		fmt.Println("Tu edad es impar")
	}
}

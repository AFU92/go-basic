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

	if edad%2 == 0 {
		fmt.Println("Tu edad es par")
	} else {
		fmt.Println("Tu edad es impar")
	}

	// Bucle for

	fmt.Println("\nEvaluar si es para cada numero del 1 a tu edad\n" + "\n.............")

	for i := 1; i <= edad; i++ {

		if i%2 == 0 {
			fmt.Println("El numero es par", i)
		} else {
			fmt.Println("El numero es impar", i)
		}
	}

	fmt.Println("\nImprimir segun elementos del slice o array\n" + "\n.............")

	array_test := []string{"elemento 1", "elemento 2", "elemento 3", "elemento 4"}

	for i := 0; i < len(array_test); i++ {

		fmt.Println("Es "+array_test[i]+" Impresion numero ", i+1)
	}

	// For each

	for _, array := range array_test {
		fmt.Println(array)
	}
}

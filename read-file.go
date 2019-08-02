package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Lector")

	new_text := []byte(os.Args[1])

	// Escribe o "reescribe el archivo"
	write := ioutil.WriteFile("file_test.txt", new_text, 0777)
	showError(write)

	fmt.Println("\nParametro recibido (Valor a insertar)")

	fmt.Println(string(new_text))

	file, errorFile := ioutil.ReadFile("file_test.txt")

	showError(errorFile)

	fmt.Println("\nArchivo en formato numerico")

	fmt.Println(file)

	fmt.Println("\nConvertirlo a string")

	fmt.Println(string(file))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}

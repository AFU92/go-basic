package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("Lector")

	file, errorFile := ioutil.ReadFile("file_test.txt")

	showError(errorFile)

	fmt.Println("Archivo en formato numerico")

	fmt.Println(file)

	fmt.Println("Convertirlo a string")

	fmt.Println(string(file))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}

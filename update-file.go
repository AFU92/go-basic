package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Lector")

	new_text := os.Args[1]

	// Escribe o "reescribe el archivo"
	// write := ioutil.WriteFile("file_test.txt", new_text, 0777)
	// showError(write)

	file, errorF := os.OpenFile("file_test.txt", os.O_APPEND, 0777)
	showError(errorF)

	writeFile, errorF := file.WriteString(new_text)
	fmt.Println(writeFile)
	showError(errorF)

	file.Close()

	text, errorFile := ioutil.ReadFile("file_test.txt")
	showError(errorFile)
	fmt.Println(string(text))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Lector de archivos")
	nuevo_texto := os.Args[1]

	escribiArchivo(nuevo_texto)

	//fmt.Println(string(nuevo_texto))s

	leerArchivo()

}

func showError(e error) {
	if e != nil {
		panic(e) //Corta ejecución y muestra error
	}
}

func leerArchivo() {
	fichero, errorArchivo := ioutil.ReadFile("fichero.txt")
	showError(errorArchivo)
	fmt.Println(string(fichero))
}

func escribiArchivo(nuevo_texto string) {

	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto + "\n")
	fmt.Println(escribir)
	showError(err)

	fichero.Close()
}

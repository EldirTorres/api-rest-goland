package main

import (
	"fmt"
	//"os" //Parametros por consola
	//"strconv"
	"time"
)

func main() {
	fmt.Println("*********** Mi programa con go *************")

	/* 	//	fmt.Println(os.Args)
	   	fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en GO")

	   	//  "_" para simplemente recibirla y no hacer nada con ella
	   	edad, _ := strconv.Atoi(os.Args[2])

	   	if (edad >= 18 || edad <= 17) && edad != 25 {
	   		fmt.Println("Eres mayor de edad")
	   	} else if edad == 25 {
	   		fmt.Println("tienes 25 años")
	   	} else if edad == 99 {
	   		fmt.Println("Pronto moriras")
	   	} else {
	   		fmt.Println("Eres menor de edad")
	   	} */

	/* numero, _ := strconv.Atoi(os.Args[1])

	if numero%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	} */

	/* 	peliculas := []string{
	"La verdad duele",
	"Ciudadano ejemplar",
	"Gran torino"} */

	//Bucle for

	/* for i := 0; i < len(peliculas); i++ {
		if i%2 == 0 {
			fmt.Println("Es par", i)
		} else {
			fmt.Println("Es impar", i)
		}
	} */

	/* for _, pelicula := range peliculas { //Como un forecah "_" para no pasarle una variables
		fmt.Println(pelicula)
	} */

	//Switch

	momento := time.Now()
	hoy := momento.Weekday()

	switch hoy {
	case 0:
		fmt.Println("Es domingo")
	case 1:
		fmt.Println("Es lunes")
	case 2:
		fmt.Println("Es martes")
	case 3:
		fmt.Println("Es miercoles")
	case 4:
		fmt.Println("Es jueves")
	case 5:
		fmt.Println("Es viernes")
	case 6:
		fmt.Println("Es sabado")
	default:
		fmt.Println("Dia no reconocido")
	}
}

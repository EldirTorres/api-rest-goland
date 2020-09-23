package main //Siempre se debe tener

import "fmt"

// fmt Mostrar cosas por pantalla hacer calculos etc

//Nota: no se puede colocar librerias y no utilizarlas

//Tipos estructurados
type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

//funcion principal
func main() {

	/* var gorra_negra = Gorra{
	marca:  "Nike",
	color:  "Negro",
	precio: 25.50,
	plana:  false}  */ //asi no da error

	/* o asi */
	/* var gorra_negra = Gorra{
	"Nike",
	"Negro",
	25.50,
	false} */

	//time.Sleep(time.Second * 5) //Para dormir el programa
	/* 	var suma int = 8 + 9
	   	var resta int = 6 - 4
	   	var nombre string = "Eldir"

	   	var apellidos string = "Torres"
	   	apellidos = "Noriega" */

	//Infiere el tipo de dato
	/* pais := "España"
	var prueba float32 = 12.3 */

	//Tipos de datos
	/*
		int
		string
		bool
		float32
	*/

	//constantes, decir no puede cambiar en el tiempo
	/* 	const year int = 2020 */

	//Operadores aritmeticos

	/* fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println("Hola mundo desde go " + nombre + apellidos + pais)
	fmt.Println("float ", prueba)
	fmt.Println("constante ", year) */

	/* var numero1 float32 = 10
	var numero2 float32 = 6 */
	/*
		fmt.Println("La suma es: ")
		fmt.Println(numero1 + numero2)
		fmt.Println("La resta es: ")
		fmt.Println(numero1 - numero2)
		fmt.Println("La multiplicación es: ")
		fmt.Println(numero1 * numero2)
		fmt.Println("La divición es: ")
		fmt.Println(numero1 / numero2)
		fmt.Println("Gorra: ")
		fmt.Println(gorra_negra)
		fmt.Println("Imprimiendo solo un valor ", gorra_negra.marca) */
	/* calculadora(numero1, numero2)
	holaMundo() */
	/* fmt.Println(devolvertexto()) */
	/* fmt.Println(gorras(45, "CLP")) */
	/* pantalones("rojo", "largo", "bolsillos", "nike") */

	//Arrays
	/* var peliculas [3]string
	peliculas[0] = "La verdad duele"
	peliculas[1] = "Ciudadano ejemplar"
	peliculas[2] = "Gran torino"

	fmt.Println(peliculas)
	fmt.Println(peliculas[1]) */

	/* peliculas := [3]string{
		"La verdad duele",
		"Ciudadano ejemplar",
		"Gran torino"}
	fmt.Println(peliculas)
	fmt.Println(peliculas[1]) */

	//Array multidimencionales
	/* var peliculas [3][2]string
	peliculas[0][0] = "La verdad duele"
	peliculas[0][1] = "Mientras duermes"

	peliculas[1][0] = "Ciudadano ejemplar"
	peliculas[1][1] = "el señor de los anillos"

	peliculas[2][0] = "Gran torino"
	peliculas[2][1] = "Gran torino"

	fmt.Println(peliculas)
	fmt.Println(peliculas[1][0]) */

	//slice, para que tengo indefinida cantidad de pociciones
	peliculas := []string{
		"La verdad duele",
		"Ciudadano ejemplar",
		"Gran torino",
		"Batman"}
	//Añadir elementos al array
	peliculas = append(peliculas, "Sin limites")

	//Saber la longitud del array
	fmt.Println("longitud del array ", len(peliculas))
	//Caracteristica, que me muestre del 0 al 3 el contenido de los indices del array
	fmt.Println("Valor de los indices de 0 al 3 ", peliculas[0:3])

	fmt.Println(peliculas)
	fmt.Println(peliculas[1])
}

//parametros
func pantalones(caracteristicas ...string) { //para que reciba variass
	for _, caracteristica := range caracteristicas { //Como un forecah "_" para no pasarle una variables
		fmt.Println(caracteristica)
	}
}

//closure
func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}
	return "El precio de gorra es ", precio(), moneda
}

func holaMundo() {
	fmt.Println("Hola mundo!! ")
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32

	if op == "+" {
		resultado = n1 + n2
	}
	if op == "-" {
		resultado = n1 - n2
	}
	if op == "*" {
		resultado = n1 * n2
	}
	if op == "/" {
		resultado = n1 / n2
	}

	return resultado
	/* fmt.Println("La resta es: ")
	fmt.Println(numero1 - numero2)
	fmt.Println("La multiplicación es: ")
	fmt.Println(numero1 * numero2)
	fmt.Println("La divición es: ")
	fmt.Println(numero1 / numero2) */
}

func calculadora(numero1 float32, numero2 float32) {
	fmt.Println("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))
	fmt.Println("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))
	fmt.Println("La multiplicación es: ")
	fmt.Println(operacion(numero1, numero2, "*"))
	fmt.Println("La divición es: ")
	fmt.Println(operacion(numero1, numero2, "/"))
}

/* func devolvertexto() (string, int) {
	dato1 := "Eldir"
	dato2 := 100
	return dato1, dato2
} */

//o
func devolvertexto() (dato1 string, dato2 int) {
	dato1 = "Eldir"
	dato2 = 100
	return
}

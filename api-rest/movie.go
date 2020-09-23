//para decirle a que package pertenece
package main

type Movie struct {
	Name     string `json:"name"` //de esta manera cambiamos el nombre
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie

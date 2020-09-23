package main //package al que pertenece

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman", 1999, "Desconocido"},
	Movie{"a todo gas", 2005, "Desconocido"},
}

/**
*Conexion a DB y se especifica la collection que usara
 */
var collection = getSession().DB("curso_go").C("movies")

/**
*Conexión a DB
 */
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

/**
*  Metodo encargado de generar las respuestas de un objeto
 */
func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

/**
*Metodo encargado de retornar la respuesta de una lista de objetos
 */
func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

/**
* Metodo de comprobación de estado del ms
 */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servicio go corriendo")
}

/**
* Metodo de retornar el listado de peliculas
 */
func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(w, 200, results)
	//json.NewEncoder(w).Encode(movies) //"w" para que escriba
	//fmt.Fprintf(w, "Listo de peliculas")
}

/**
* Metodo que retorna la pelicula
 */
func MovieShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) //Recogemos variable de url
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) { //comprueba si es un hexadecimal
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)

	//fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id) // -> "%s" lo sustituye por la variable
}

/**
* Metodo que añade una pelicula
 */
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body) //Decodifica el json

	var movie_data Movie
	err := decoder.Decode(&movie_data) //"&" para indicar que es una variable que aun no tiene nada

	if err != nil {
		panic(err)
	}

	defer r.Body.Close() //"defer" cierra la lectura de algo

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	//log.Println(movie_data)             //Imprimir por consola
	//movies = append(movies, movie_data) //Añadir elementos a un array
	responseMovie(w, 200, movie_data)
}

/**
* Metodo que actualiza una pelicula
 */
func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}       //ubica el id
	change := bson.M{"$set": movie_data} //setea la data
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

/**
*Metodos de acceso para setear datos
 */
func (this *Message) setStatus(data string) { //"*" indica que apuanta a esa dirección en memoria
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

/**
* Metodo que elimina una pelicula
 */
func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//results := Message{"success", "La pelicula con ID "+movie_id+" ha sido borrada correctamente"}
	message := new(Message) //Cuando se crea un objeto se debe especificar con * que apunte a esa dirección en memoria EJ: linea 184 y 188

	message.setStatus("success")
	message.setMessage("La pelicula con ID " + movie_id + " ha sido borrada correctamente")

	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

# api-rest-goland
Proyecto que contempla el desarrollo de un api rest bajo goland, y ejemplos de como funciona go desde un punto de vista general

#Pasos para correr el proyecto

1) Instalar go -> https://golang.org/
1.1) Instalar mongoDB -> https://www.mongodb.com/
2) Intalar los siguientes paquetes:
    -   rutas -> go get -u github.com/gorilla/mux
    -> plugin conexion mongo : nos ubicamos en la carpeta de nuestro proyecto
        -> https://labix.org/mgo
            -> API docs for mgo   
                -> go get gopkg.in/mgo.v2

        -> plugin para trabajar con bson
            -> API docs for mgo/bson
                -> go get gopkg.in/mgo.v2/bson

3) para correr o leventar el proyecto:
    situarnos en el directorio api-rest
                        -> go run *.go


Nota: si se quiere probar los ejemplos que menciono deben situarce en el archivo y hacer lo siguiente:

    Ejemplo: go run hola-mundo.go
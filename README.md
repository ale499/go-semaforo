# go-semaforo
    VERSION ESTABLE.
Descripcion del proyecto:
La creación de semaforos inteligentes mostrando sus estados en tiempo real simulando el sistema de transito en todo sus estados (verde, amarillo y rojo) para esto hemos utilizado el lenguaje desarrollado por Google llamado Go, y sus frameworks Gin, Gorm y Compile Daemons. Para realizar el proyecto decidimos crear una base de datos con residencia en la nube utilizando Elephant SQL y TablePlus, todo esto implementado en el IDE Visual Studio Code con la version 1.20 de Go.

Mediante la extension de thunder client nos vamos a mover para poder ver/agregar/borrar/editar los semaforos y mostrarlos en el postman (La consola de la parte derecha) con los comandos que estan señalados en el main y colocar el url que es http://localhost:3000/accion(post,get,put,delete)/(id del semaforo 1-15)/
POST("/crear", controllers.Postsemaforo)
 GET("/getall", controllers.Getsemaforo)
GET("/traer/:id", controllers.Getidsemaforo) 
PUT("/actualizar/:id", controllers.Updatesemaforo) 
DELETE("/borrar/:id", controllers.Deletesemaforo)
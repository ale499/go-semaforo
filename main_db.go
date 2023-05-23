package main

import (
	"github.com/TadeLauda/go-semaforo/controllers"
	"github.com/TadeLauda/go-semaforo/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.Conectdb()
}

func main() {
	r := gin.Default()
	r.POST("/crear", controllers.Postsemaforo)
	r.GET("/getall", controllers.Getsemaforo)
	r.GET("/traer/:id", controllers.Getidsemaforo)      //http://localhost:3000/traer/2//
	r.PUT("/actualizar", controllers.Updatesemaforo)    //TODO: falta terminar este metodo aplicando paralelismo//
	r.DELETE("/borrar/:id", controllers.Deletesemaforo) //http://localhost:3000/borrar/1//

	r.Run()
}

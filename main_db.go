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
	r.GET("/traer/:id", controllers.Getidsemaforo)
	r.PUT("/actualizar", controllers.Updatesemaforo)
	r.DELETE("/borrar", controllers.Deletesemaforo)
	r.Run()
}

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
	r.Run() // listen and serve on 0.0.0.0:8080
}

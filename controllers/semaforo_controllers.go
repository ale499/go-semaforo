package controllers

import (
	"github.com/TadeLauda/go-semaforo/initializers"
	"github.com/TadeLauda/go-semaforo/models"
	"github.com/gin-gonic/gin"
)

func Postsemaforo(c *gin.Context) {
	//vamos a obtener los datos del body
	var body struct {
		Estado    string
		Ubicacion string
	}

	c.Bind(&body)
	//crear semaforo

	semaforo := models.Semaforo{Estado: body.Estado, Ubicacion: body.Ubicacion}
	resultado := initializers.Db.Create(&semaforo)
	if resultado.Error != nil {

		c.Status(400)

		return

	}
	//retornar semaforo

	c.JSON(200, gin.H{
		"semaforo": semaforo,
	})

}

func Getsemaforo(c *gin.Context) {
	/* obtener semaforos en un slice*/
	var semaforos []models.Semaforo
	initializers.Db.Find(&semaforos)

	// retornar semaforos //
	c.JSON(200, gin.H{
		"Todos los semaforos": semaforos,
	})
}

func Getidsemaforo(c *gin.Context) {
	//obtener id del semaforo//
	id := c.Param("id")
	// obtener el semaforo//
	var semaforo models.Semaforo
	initializers.Db.First(&semaforo, id)
	//retornar semaforo//
	c.JSON(200, gin.H{
		"semaforo": semaforo,
	})
}

func Updatesemaforo(c *gin.Context) {
	// TODO: usar este metodo combinandolo con paralelismo para poder manejar el estado de los semaforos en tiempo real
}
func Deletesemaforo(c *gin.Context) {
	//Obtener id del semaforo
	id := c.Param("id")
	//Borrar el semaforo
	initializers.Db.Delete(&models.Semaforo{}, id)
	//Responder
	c.Status(200)

}

package controllers

import (
	"time"

	"github.com/TadeLauda/go-semaforo/initializers"
	"github.com/TadeLauda/go-semaforo/logic"
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
	time.Sleep(10)
	//obtener id del semaforo//
	id := c.Param("id")
	// obtener el semaforo//
	var semaforo models.Semaforo
	initializers.Db.First(&semaforo, id)
	//retornar semaforo//
	c.JSON(200, gin.H{
		"semaforo": semaforo,
	})
	//Cada vez que se haga una peticion se ejecutara una GO ROUTINE en paralelo para determinar el cambio de los semaforos en función de su estado anterior.

	go logic.ChangelogicState(semaforo.Estado)

}
func Updatesemaforo(c *gin.Context) {

	//Obtenemos el id del semaforo a modificar
	id := c.Param("id")
	//Obtenemos los datos del semaforo
	var body struct {
		Estado    string
		Ubicacion string
	}
	c.Bind(&body)

	//Encontramos el semaforo que queremos actualizar
	var semaforo models.Semaforo
	initializers.Db.First(&semaforo, id)

	newState := logic.ChangelogicState(semaforo.Estado)
	//Lo actualizamos
	initializers.Db.Model(&semaforo).Updates(models.Semaforo{
		Estado:    newState,
		Ubicacion: body.Ubicacion,
	})

	c.JSON(200, gin.H{
		"Semaforo": semaforo,
	})
}
func Deletesemaforo(c *gin.Context) {
	//Obtener id del semaforo
	id := c.Param("id")
	//Borrar el semaforo
	initializers.Db.Delete(&models.Semaforo{}, id)
	//Responder
	c.Status(200)
}

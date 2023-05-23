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
func Getsemaforo(c *gin.Context)    {}
func Getidsemaforo(c *gin.Context)  {}
func Updatesemaforo(c *gin.Context) {}
func Deletesemaforo(c *gin.Context) {}

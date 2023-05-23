package main

import (
	"github.com/TadeLauda/go-semaforo/initializers"
	"github.com/TadeLauda/go-semaforo/models"
)

func init() {
	initializers.Loadenv()
	initializers.Conectdb()
}

func main() {
	initializers.Db.AutoMigrate(&models.Semaforo{})
}

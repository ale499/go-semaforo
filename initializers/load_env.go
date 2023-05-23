package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando las variables de entorno")
	}
}

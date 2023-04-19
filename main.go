package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Comienzo:")
	fmt.Println("PRIMERA LINEA DE COMANDO PARA HACER EL PROYECTO")

	fmt.Println("Primera Prueba")

	fmt.Println("Funcionamiento basico de semaforos: ")

	for {
		// Estado: Rojo
		fmt.Println("Semaforo: Rojo")
		time.Sleep(10 * time.Second) //Espera 10 segundos

		// Estado: Amarillo
		fmt.Println("Semaforo: Amarillo")

		// Estado: Verde
		time.Sleep(2 * time.Second) //Espera 2 segundos
		fmt.Println("Semaforo: Verde")
		time.Sleep(12 * time.Second) //Espera 12 segundos
	}

}

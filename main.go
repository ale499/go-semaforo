package main

import (
	"fmt"
	"time"
)

func main() {

	var n bool

	fmt.Println("Comienzo:")
	fmt.Println("PRIMERA LINEA DE COMANDO PARA HACER EL PROYECTO")

	fmt.Println("Segunda Prueba")

	fmt.Println("Funcionamiento basico de semaforos: ")

	fmt.Print("Es una calle atrevesada por otra?")

	fmt.Scan(&n)

	fmt.Println("Prueba linux lauti")

	fmt.Println("Funciono en la maquna virtual de Tadeo")

	fmt.Println("Funciono en la maquina virtual de Lucas")

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

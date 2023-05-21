package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaforo struct {
	color    string
	duracion time.Duration
}

func cambiarSemaforo(semaforo *Semaforo, wg *sync.WaitGroup, ch chan *Semaforo) {
	defer wg.Done()

	for {
		fmt.Printf("%s semaforo encendido\n", semaforo.color)
		time.Sleep(semaforo.duracion)

		fmt.Printf("%s semaforo apagado\n", semaforo.color)
		time.Sleep(time.Second)

		ch <- semaforo // Enviar semáforo actual al canal
		<-ch           //Esperar para recibir el siguiente semaforo
	}
}

// Esta funcion recibe y cambia los semaforos en el orden correcto.
// Utiliza ch para recibir y enviar los semaforos.
// Utilizamos las go routines tambien para llamar a la funcion cambiar semaforo
func controladorSemaforos(semaforos []*Semaforo) {
	ch := make(chan *Semaforo)

	var wg sync.WaitGroup
	wg.Add(len(semaforos))

	for _, semaforo := range semaforos {
		go cambiarSemaforo(semaforo, &wg, ch)
	}

	wg.Wait()
}
func main() {

	//var n bool

	fmt.Println("Comienzo:")
	fmt.Println("PRIMERA LINEA DE COMANDO PARA HACER EL PROYECTO")

	fmt.Println("Segunda Prueba")

	fmt.Println("Funcionamiento basico de semaforos: ")

	//fmt.Print("Es una calle atrevesada por otra?")

	//fmt.Scan(&n)

	//for {
	// Estado: Rojo
	//fmt.Println("Semaforo: Rojo")
	//time.Sleep(10 * time.Second) //Espera 10 segundos

	// Estado: Amarillo
	//fmt.Println("Semaforo: Amarillo")

	// Estado: Verde
	//time.Sleep(2 * time.Second) //Espera 2 segundos
	//fmt.Println("Semaforo: Verde")
	//time.Sleep(12 * time.Second) //Espera 12 segundos
	//}

	fmt.Println("Cambio desde Windows a linux de Alejandro")

	fmt.Println("Prueba linux lauti")

	fmt.Println("Funciono en la maquna virtual de Tadeo")

	fmt.Println("Cambio de entorno a la maquina virtual de Lucas")

	//Prueba preliminar semaforo (Prueba muuuuy simple)
	//Creo 3 semaforos de cada color y con distinta duracion, lo almaceno en la lista
	//semaforos y luego lo paso a la funcion controlador semaforos

	semaforo1 := &Semaforo{color: "Rojo", duracion: 5 * time.Second}
	semaforo2 := &Semaforo{color: "Amarillo", duracion: 2 * time.Second}
	semaforo3 := &Semaforo{color: "Verde", duracion: 3 * time.Second}

	semaforos := []*Semaforo{semaforo1, semaforo2, semaforo3}

	controladorSemaforos(semaforos)
}

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

func cambiarSemaforo(semaforo *Semaforo, ch chan<- *Semaforo) {
	fmt.Printf("%s semaforo encendido\n", semaforo.color)
	time.Sleep(semaforo.duracion)

	fmt.Printf("%s semaforo apagado\n", semaforo.color)
	time.Sleep(time.Second)

	ch <- semaforo // Enviar semáforo actual al canal
}

// Esta funcion recibe y cambia los semaforos en el orden correcto.
// Utiliza ch para recibir y enviar los semaforos.
// Utilizamos las go routines tambien para llamar a la funcion cambiar semaforo
func controladorSemaforos(semaforos []*Semaforo, ch <-chan *Semaforo, cambio <-chan struct{}) {
	var siguiente *Semaforo

	for {
		semaforo := <-ch // Recibir semáforo actual del canal

		// Encontrar el siguiente semáforo en la lista
		for i, s := range semaforos {
			if s == semaforo {
				siguiente = semaforos[(i+1)%len(semaforos)]
				break
			}
		}

		time.Sleep(time.Second) // Esperar antes de cambiar al siguiente semáforo

		fmt.Printf("Cambiando de %s a %s semaforo\n", semaforo.color, siguiente.color)

		<-cambio // Esperar señal de la goroutine anónima
	}
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

	ch := make(chan *Semaforo, 2)
	cambio := make(chan struct{})

	go controladorSemaforos(semaforos, ch, cambio)

	var wg sync.WaitGroup
	wg.Add(len(semaforos))

	for _, semaforo := range semaforos {
		go func(s *Semaforo) {
			cambiarSemaforo(s, ch)
			cambio <- struct{}{} //Señalar cambio de semaforo a la goroutine controladorSemaforos
			wg.Done()
		}(semaforo)
	}
	wg.Wait()
}

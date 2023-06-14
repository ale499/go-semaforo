package logic

import "time"

func ChangelogicState(estado string) string {
	//Logica para cambiar el color del semaforo
	//Los timer sirven para que despues de una cantidad determinada de tiempo se cambie el color en función de su estado anterior
	if estado == "rojo" {
		time.After(1)
		estado = "verde"
	} else if estado == "amarillo" {
		time.After(1)
		estado = "rojo"
	} else if estado == "verde" {
		time.After(1)
		estado = "amarillo"
	}
	return estado
}

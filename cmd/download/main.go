package main

import (
	"fmt"
	"golang-dev-challenge-mid/pkg/download"
)

func main() {

	if err := download.Run("http://bibliotecadigital.ilce.edu.mx/Colecciones/ObrasClasicas/_docs/Corazon_Amicis.pdf"); err != nil {
		fmt.Println("ERROR: ", err)
	}
}

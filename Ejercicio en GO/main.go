package main

import "fmt"

func main() {

	/*var numero2 int64

	numero := 0

	fmt.Println("indique su nombre")
	fmt.Scanln(&numero2)
	fmt.Println("Indique su edad")
	fmt.Scanln(&numero)*/

	var Kilometro, Millas float64 = 0.0, 0.0
	fmt.Println("Converso de Kilometros A Millas")
	fmt.Println("Ingrese Kilometros")
	fmt.Scanln(&Kilometro)

	Millas = Kilometro * 0.621
	fmt.Println("la conversion de", Kilometro, "Km es de", Millas, "M")
}

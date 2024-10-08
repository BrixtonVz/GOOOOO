package main

import (
    "fmt"
    "strconv"
)

func Infanteria(Reclutas int, TiempoEnCombate int) int {
    Experiencia := 0.05 // 5% de incremento de infanteria por año de servicio
    InfanteriaTotal := int64(Reclutas) * (1 + Experiencia*float64(TiempoEnCombate))
    return int(InfanteriaTotal)
}

func main() {
    // Pedimos al usuario que ingrese la cantidad de reclutas
    fmt.Println("Ingrese la cantidad de reclutas:")
    var inputReclutas string
    fmt.Scanln(&inputReclutas)

    // Convertimos la entrada de los reclutas a un entero
    Reclutas, err := strconv.Atoi(inputReclutas)
    if err != nil {
        fmt.Println("Error al convertir los Reclutas a int:", err)
        return
    }

    // Pedimos al usuario que ingrese El tiempo que lleva en combate
    fmt.Println("Ingresa tus años de servicio:")
    var inputAños string
    fmt.Scanln(&inputAños)

    // Convertimos la entrada de años en combate de servicio a un entero
    añosServicio, err := strconv.Atoi(inputAños)
    if err != nil {
        fmt.Println("Error al convertir los años de servicio a int:", err)
        return
    }

    // Calculamos la nueva cantidad de reclutas
    Infanteria := Infanteria(Reclutas, TiempoEnCombate)

    // Mostramos la cantidad total de nuevos reclutas o infanteria
    fmt.Printf("la nueva cantidad de infanteria otorgada es de: %d\n",Infanteria)
}
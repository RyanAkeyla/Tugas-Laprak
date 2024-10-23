
package main

import "fmt"

func main() {
    var Celcius float64
    fmt.Print("Masukkan suhu dalam Celcius: ")
    fmt.Scanln(&Celcius)
    Fahrenheit := (Celcius * 9 / 5) + 32
    fmt.Println("Suhu dalam Fahrenheit: ", Fahrenheit, )
    Reamur := Celcius * 4 / 5
    fmt.Println("Suhu dalam Reamur: ", Reamur, )
    Kelvin := Celcius + 273.15
    fmt.Println("Suhu dalam Kelvin: ", Kelvin, )
}

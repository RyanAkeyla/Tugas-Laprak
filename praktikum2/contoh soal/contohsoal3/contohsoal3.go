
package main

import "fmt"

func main() {
	var Rupiah, Dollar float64
	fmt.Print("Masukkan Nominal Rupiah :")
	fmt.Scan(&Rupiah)
	Dollar = (Rupiah / 15000)
	fmt.Print("jadi", Rupiah, "Rupiah=", Dollar, "Dollar")
}

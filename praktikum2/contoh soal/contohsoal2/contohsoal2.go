
package main

import "fmt"

func main() {
	var Alas, Tinggi, Luas float64
	fmt.Print("Masukkan Alas :")
	fmt.Scan(&Alas)
	fmt.Print("Tinggi :")
	fmt.Scan(&Tinggi)
	Luas = (Alas * Tinggi /2)
	fmt.Println(Luas)
}


package main

import "fmt"

func main() {
	var t int
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&t)
	kabisat := (t%400 == 0) || (t%4 == 0 && t%100 != 0)
	fmt.Print(kabisat)
}

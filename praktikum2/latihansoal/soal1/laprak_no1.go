
package main

import "fmt"

func main() {
	var fx float64
	fmt.Println("Masukkan nilai f(x): ")
	fmt.Scan(&fx)
	x := 2/(fx+5) + 5
	fmt.Println("Nilai x adalah:", x)
}

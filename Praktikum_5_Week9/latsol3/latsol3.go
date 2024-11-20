package main

import "fmt"

func main() {
	var x, y int
	var hasil1, hasil2 bool
	fmt.Scan(&x, &y)
	hasil2 = x%y == 0
	if y%x == 0 {
		hasil1 = true
	}
	fmt.Println(hasil1, hasil2)
}
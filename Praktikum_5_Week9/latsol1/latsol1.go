package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	jumlahMotor := n / 2
	if n%2 != 0 {
		jumlahMotor += 1
	}
	fmt.Println(jumlahMotor)
}

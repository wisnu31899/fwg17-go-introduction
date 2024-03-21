package main

import (
	"fmt"
)

// menggunakan loop dan condition
func printSegitiga(n int) {
	// Mencetak baris
	for i := n; i >= 1; i-- {
		// Mencetak spasi di awal setiap baris
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// Mencetak bintang
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

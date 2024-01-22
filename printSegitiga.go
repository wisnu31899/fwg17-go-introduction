package main

import "fmt"

//fix
//output :
// 5 5 5 5 5
//  4 4 4 4
//   3 3 3
//    2 2
//     1
//menggunakan loop and condition
func printSegitiga(n int) {
	for i := n; i >= 1; i-- {
		// Mencetak spasi di awal setiap baris
		for j := 0; j < n; j++ {
			if j < n-i {
				fmt.Print(" ")
			} else {
				// Mencetak angka dari i hingga 1
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	}
}

//menggunakan loop
// func printSegitiga(n int) {
// 	for i := n; i >= 1; i-- {
// 		// Mencetak spasi di awal setiap baris
// 		for j := 0; j < n-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		// Mencetak angka dari i hingga 1
// 		for k := 1; k <= i; k++ {
// 			fmt.Print(i, " ")
// 		}
// 		fmt.Println()
// 	}
// }

//output :
// 55555
//  4444
//   333
//    22
//     1
// func printSegitiga(n int) {
// 	for i := n; i >= 1; i-- {
// 		// Mencetak spasi di awal setiap baris
// 		for j := 0; j < n-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		// Mencetak angka dari i hingga 1
// 		for k := 1; k <= i; k++ {
// 			fmt.Print(i)
// 		}
// 		// Mencetak angka dari 2 hingga i
// 		// for l := 2; l <= i; l++ {
// 		// 	fmt.Print(l)
// 		// }
// 		fmt.Println()
// 	}
// }

//output:
// 54321
//  4321
//   321
//    21
//     1
// func printSegitiga(n int) {
// 	for i := n; i >= 1; i-- {
// 		// Mencetak spasi di awal setiap baris
// 		for j := 0; j < n-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		// Mencetak angka dari i hingga 1
// 		for k := i; k >= 1; k-- {
// 			fmt.Print(k)
// 		}
// 		fmt.Println()
// 	}
// }

// output:
// 5 4 3 2 1
//  4 3 2 1
//   3 2 1
//    2 1
//     1
// func printSegitiga(n int) {
// 	for i := n; i >= 1; i-- {
// 		// Mencetak spasi di awal setiap baris
// 		for j := 0; j < n-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		// Mencetak angka dari i hingga 1
// 		for k := i; k >= 1; k-- {
// 			fmt.Print(k, " ")
// 		}
// 		fmt.Println()
// 	}
// }
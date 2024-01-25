package main

import "fmt"

// func main() {
// 	fmt.Println("hello word")
// }

// //soal 1
// func main() {
// 	// Contoh penggunaan
// 	printSegitiga(5)
// }

// soal 2
func main() {
	// Contoh penggunaan (low, medium, strong)
	password := generatePassword("wisnuhidayat", "strong")
	fmt.Println("Hasil generate password:", password)
}

//soal 3
// func main() {
// 	data := []int{1, 7, 3, 4, 8, 9}
// 	n := 40
// 	result1, result2 := findMovieDuration(data, n)
// 	if result1 != -1 && result2 != -1 {
// 		fmt.Printf("Dua durasi film yang hasil penjumlahannya sama dengan %d jam perjalanan adalah %d jam dan %d jam\n", n, result1, result2)
// 	} else {
// 		fmt.Printf("Tidak ada dua durasi film yang hasil penjumlahannya sama dengan %d jam perjalanan\n", n)
// 	}
// }

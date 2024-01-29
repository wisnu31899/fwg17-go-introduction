package main

import (
	"math/rand"
)

// func generatePassword(password, level string) string {

// 	// Memastikan panjang password minimal 6 karakter
// 	if len(password) < 6 {
// 		return "Password length should be at least 6 characters."
// 	}

// 	// Fungsi untuk mendapatkan karakter acak dari setiap jenis karakter
// 	getRandomLower := func() byte {
// 		randomLower := "abcdefghijklmnopqrstuvwxyz"
// 		return randomLower[rand.Intn(len(randomLower))]
// 	}
// 	getRandomUpper := func() byte {
// 		RandomUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 		return RandomUpper[rand.Intn(len(RandomUpper))]
// 	}
// 	getRandomDigit := func() byte {
// 		RandomDigit := "0123456789"
// 		return RandomDigit[rand.Intn(len(RandomDigit))]
// 	}
// 	getRandomSpecial := func() byte {
// 		specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
// 		return specialChars[rand.Intn(len(specialChars))]
// 	}

// 	var newPassword []byte //buatkan var dgn type slice byte

// 	for i := 0; i < len(password); i++ {
// 		if level == "low" {
// 			randType := rand.Intn(2)
// 			if randType == 0 {
// 				newPassword = append(newPassword, getRandomLower())
// 			} else if randType == 1 {
// 				newPassword = append(newPassword, getRandomUpper())
// 			}
// 		} else if level == "medium" {
// 			randType := rand.Intn(3)
// 			if randType == 0 {
// 				newPassword = append(newPassword, getRandomLower())
// 			} else if randType == 1 {
// 				newPassword = append(newPassword, getRandomUpper())
// 			} else if randType == 2 {
// 				newPassword = append(newPassword, getRandomDigit())
// 			}
// 		} else if level == "strong" {
// 			randType := rand.Intn(4)
// 			if randType == 0 {
// 				newPassword = append(newPassword, getRandomLower())
// 			} else if randType == 1 {
// 				newPassword = append(newPassword, getRandomUpper())
// 			} else if randType == 2 {
// 				newPassword = append(newPassword, getRandomDigit())
// 			} else if randType == 3 {
// 				newPassword = append(newPassword, getRandomSpecial())
// 			}
// 		}
// 	}

// 	return string(newPassword)
// }

func generatePassword(password, level string) string {
	// Memastikan panjang password minimal 6 karakter
	if len(password) < 6 {
		return "Password length should be at least 6 characters."
	}

	// Fungsi untuk mendapatkan karakter acak dari setiap jenis karakter
	getRandomLower := func() string {
		randomLower := "abcdefghijklmnopqrstuvwxyz"
		return string(randomLower[rand.Intn(len(randomLower))])
	}
	getRandomUpper := func() string {
		RandomUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return string(RandomUpper[rand.Intn(len(RandomUpper))])
	}
	getRandomDigit := func() string {
		RandomDigit := "0123456789"
		return string(RandomDigit[rand.Intn(len(RandomDigit))])
	}
	getRandomSpecial := func() string {
		specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
		return string(specialChars[rand.Intn(len(specialChars))])
	}

	var newPassword string // ubah tipe data ke string

	for i := 0; i < len(password); i++ {
		if level == "low" {
			randType := rand.Intn(2)
			if randType == 0 {
				newPassword += getRandomLower()
			} else if randType == 1 {
				newPassword += getRandomUpper()
			}
		} else if level == "medium" {
			randType := rand.Intn(3)
			if randType == 0 {
				newPassword += getRandomLower()
			} else if randType == 1 {
				newPassword += getRandomUpper()
			} else if randType == 2 {
				newPassword += getRandomDigit()
			}
		} else if level == "strong" {
			randType := rand.Intn(4)
			if randType == 0 {
				newPassword += getRandomLower()
			} else if randType == 1 {
				newPassword += getRandomUpper()
			} else if randType == 2 {
				newPassword += getRandomDigit()
			} else if randType == 3 {
				newPassword += getRandomSpecial()
			}
		}
	}

	return newPassword
}

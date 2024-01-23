package main

import (
	"math/rand"
)

func generatePassword(password, level string) string {

	// Fungsi untuk mendapatkan karakter acak dari setiap jenis karakter
	getRandomLower := func() byte {
		return byte(rand.Intn(26) + 'a')
	}
	getRandomUpper := func() byte {
		return byte(rand.Intn(26) + 'A')
	}
	getRandomDigit := func() byte {
		return byte(rand.Intn(10) + '0')
	}
	getRandomSpecial := func() byte {
		specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
		return specialChars[rand.Intn(len(specialChars))]
	}

	// Fungsi untuk memanipulasi password sesuai level
	manipulatePassword := func(pwd string) string {
		var newPassword []byte

		for i := 0; i < len(pwd); i++ {
			if level == "low" {
				randType := rand.Intn(2)
				if randType == 0 {
					newPassword = append(newPassword, getRandomLower())
				} else if randType == 1 {
					newPassword = append(newPassword, getRandomUpper())
				}
			} else if level == "medium" {
				randType := rand.Intn(3)
				if randType == 0 {
					newPassword = append(newPassword, getRandomLower())
				} else if randType == 1 {
					newPassword = append(newPassword, getRandomUpper())
				} else if randType == 2 {
					newPassword = append(newPassword, getRandomDigit())
				}
			} else if level == "strong" {
				randType := rand.Intn(4)
				if randType == 0 {
					newPassword = append(newPassword, getRandomLower())
				} else if randType == 1 {
					newPassword = append(newPassword, getRandomUpper())
				} else if randType == 2 {
					newPassword = append(newPassword, getRandomDigit())
				} else if randType == 3 {
					newPassword = append(newPassword, getRandomSpecial())
				}
			}
		}

		return string(newPassword)
	}

	// Memastikan panjang password minimal 6 karakter
	if len(password) < 6 {
		return "Password length should be at least 6 characters."
	}

	// Memanipulasi password sesuai dengan level
	newPassword := manipulatePassword(password)

	return newPassword
}

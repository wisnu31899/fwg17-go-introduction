package main

//n lama perjalanan, data slice waktu
func findMovieDuration(data []int, n int) (int, int) {
	movieMap := make(map[int]int)

	for i, duration := range data {
		duration2 := n - duration
		if index, found := movieMap[duration2]; found {
			// Temukan pasangan yang cocok
			return data[index], data[i]

		}
		// Tambahkan durasi ke dalam peta
		movieMap[duration] = i
		// println(movieMap[duration])
	}

	// Jika tidak ada pasangan yang cocok
	return -1, -1
}

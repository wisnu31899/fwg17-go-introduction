package main

func findMovieDuration(data []int, n int) (int, int) {
	movieMap := make(map[int]int)

	for i, duration := range data {
		complement := n - duration
		if index, found := movieMap[complement]; found {
			// Temukan pasangan yang cocok
			return data[index], data[i]
		}
		// Tambahkan durasi ke dalam peta
		movieMap[duration] = i
	}

	// Jika tidak ada pasangan yang cocok
	return -1, -1
}

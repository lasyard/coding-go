package popcount

func PopCount3(x uint64) int {
	var count int
	for x != 0 {
		count++
		x = x & (x - 1)
	}
	return count
}

package popcount

func PopCount2(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}

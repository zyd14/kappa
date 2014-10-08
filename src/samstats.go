package sam

func Average (array [] int) float32 {
	n := len (array)
	var sum int = 0
	for i := 0; i < n; i++ {
		sum = sum + array[i]
	}
	return float32(sum/n)
}
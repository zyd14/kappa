package kstat

func average (data [] int) float64 {
	sum := data[0]
	for i:= 1; i < len (data); i++ {
		sum = sum + data [i]
	}
	var mean float64 = float64 (sum / len (data))
	return mean
}

func max (data [] int) int {
	if len (data) == 0 {
		return 0
	} else {
		max := data[0]
		for i:= 1; i < len (data); i++ {
			if data[i] > max {
				max = data[i]
			}
		}
		return max
	}
}

func min (data [] int) int {
	if len (data) == 0 {
		return 0
	} else {
		min := data[0]
		for i:= 1; i < len (data); i++ {
			if data[i] < min {
				min = data[i]
			}
		}
		return min
	}
}

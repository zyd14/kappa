package kstat

func Echo (msg string) string {
	return msg
}

func GetMax (a, b float64) float64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Average (data []float64) float64 {
	if len (data) <= 0 {
		return float64(0)
	}
	sum := data[0]
	for i:= 1; i < len (data); i++ {
		sum = sum + data [i]
	}
	var mean float64 = sum / float64 (len (data))
	return mean
}

func Max (data []float64) float64 {
	if len (data) == 0 {
		return float64(0)
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

func Min (data []float64) float64 {
	if len (data) == 0 {
		return float64(0)
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

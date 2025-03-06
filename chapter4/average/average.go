package average

func AvgFloatArray(arr []float64) float64 {
	var sum float64
	for _, val := range arr {
		sum += val
	}
	return sum / float64(len(arr))
}

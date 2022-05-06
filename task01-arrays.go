package homework

func average(input [15]float32) (result float32) {
	var count float32
	for _, val := range input {
		count += val
	}
	return count / 15
}

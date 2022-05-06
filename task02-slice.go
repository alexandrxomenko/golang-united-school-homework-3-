package homework

func reverse(input []int64) (result []int64) {

	for key := range input {
		result = append(result, input[len(input)-key-1])
	}
	return
}

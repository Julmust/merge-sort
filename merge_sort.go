package merge_sort

func sort(input []int) []int {
	var output []int
	slice_length := len(input)
	if slice_length == 1 {
		return input
	}
	split_point := slice_length / 2
	low := sort(input[:split_point])
	high := sort(input[split_point:])

	for {
		if len(low) == 0 {
			output = append(output, high...)
			break
		} else if len(high) == 0 {
			output = append(output, low...)
			break
		} else if low[0] < high[0] {
			output = append(output, low[0])
			low = low[1:]
		} else { // high > low OR high == low
			output = append(output, high[0])
			high = high[1:]
		}
	}

	return output
}

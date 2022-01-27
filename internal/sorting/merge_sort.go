package sorting

func MergeSort(inputSlice []int64) []int64 {
	if len(inputSlice) <= 1 {
		return inputSlice
	}

	left := make([]int64, 0)
	right := make([]int64, 0)

	for index, num := range inputSlice {
		if index < len(inputSlice)/2 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left []int64, right []int64) []int64 {
	res := make([]int64, 0)

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}

	return res
}

package main

func mergeSorted(arr1, arr2 []int) []int {
	var result = make([]int, len(arr1)+len(arr2))

	var i, j, k int
	for ; k < len(result); k++ {
		if i >= len(arr1) {
			result[k] = arr2[j]
			j++
			continue
		} else if j >= len(arr2) {
			result[k] = arr1[i]
			i++
			continue
		}

		if arr1[i] <= arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
	}

	return result
}

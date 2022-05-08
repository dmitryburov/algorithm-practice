package main

func merge(arr []int, left int, mid int, right int) []int {
	midCopy := mid
	result := make([]int, 0, right-left)

	for left < mid || midCopy < right {
		if left == mid {
			result = append(result, arr[midCopy])
			midCopy++
			continue
		}
		if midCopy == right {
			result = append(result, arr[left])
			left++
			continue
		}
		if arr[left] <= arr[midCopy] {
			result = append(result, arr[left])
			left++
		} else {
			result = append(result, arr[midCopy])
			midCopy++
		}
	}

	return result
}

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf <= 1 {
		return
	}

	mid := (lf + rg) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	res := merge(arr, lf, mid, rg)

	k := 0
	for i := lf; i < rg; i++ {
		arr[i] = res[k]
		k++
	}
}

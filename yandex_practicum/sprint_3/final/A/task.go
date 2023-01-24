package main

// brokenSearch основная функция
func brokenSearch(arr []int, k int) int {
	return binarySearch(arr, k, 0, len(arr)-1)
}

// binarySearch бинарный поиск по массиву
func binarySearch(arr []int, k, left, right int) int {

	if left > right {
		return -1
	}

	mid := (right + left) / 2
	if arr[mid] == k {
		return mid
	}

	// если левая часть отсортирована
	if arr[left] <= arr[mid] {
		if k >= arr[left] && k <= arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		// тогда по правой стороне
	} else {
		if k >= arr[mid] && k <= arr[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// рекурсивно идем дальше с учетём измененного отрезка
	return binarySearch(arr, k, left, right)
}

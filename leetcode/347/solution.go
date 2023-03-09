package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	var hMap = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		hMap[nums[i]]++
	}

	keys := make([]int, len(hMap))
	for key := range hMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(a, b int) bool {
		return hMap[keys[a]] > hMap[keys[b]]
	})

	var result = make([]int, k)
	for i := 0; i < len(keys) && i < k; i++ {
		result[i] = keys[i]
	}

	return result
}

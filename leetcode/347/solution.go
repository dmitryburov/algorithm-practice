package main

func topKFrequent(nums []int, k int) []int {
	var result []int

	var numMap = make(map[int]int)
	for _, val := range nums {
		numMap[val]++
	}

	freqMap := make(map[int][]int)
	for key, freq := range numMap {
		freqMap[freq] = append(freqMap[freq], key)
	}

	for i := len(nums); len(result) != k; i-- {
		for _, n := range freqMap[i] {
			if len(result) != k {
				result = append(result, n)
			}
		}
	}

	return result
}

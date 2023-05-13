package main

func twoSum(numbers []int, target int) []int {
	hMap := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		if j, ok := hMap[target-numbers[i]]; ok {
			return []int{j + 1, i + 1}
		}
		hMap[numbers[i]] = i
	}

	return nil
}

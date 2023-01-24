package main

func twoSum(nums []int, target int) []int {
	hMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := hMap[target-nums[i]]; ok {
			return []int{j, i}
		}
		hMap[nums[i]] = i
	}

	return nil
}

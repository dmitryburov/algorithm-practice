package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)

	for _, interval := range intervals {
		if len(result) == 0 || result[len(result)-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			max := result[len(result)-1][1]
			if result[len(result)-1][1] < interval[1] {
				max = interval[1]
			}
			result[len(result)-1][1] = max
		}
	}

	return result
}

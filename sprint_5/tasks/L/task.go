package main

func siftDown(heap []int, idx int) int {
	for {
		left := idx * 2
		right := idx*2 + 1

		if left >= len(heap) {
			break
		}

		max := 0

		if right < len(heap) && heap[right] > heap[left] {
			max = right
		} else {
			max = left
		}

		if heap[max] > heap[idx] {
			heap[max], heap[idx] = heap[idx], heap[max]
			idx = max
		} else {
			break
		}
	}

	return idx
}

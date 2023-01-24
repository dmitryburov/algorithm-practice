package main

func siftUp(heap []int, idx int) int {

	for heap[idx] > heap[idx/2] {
		if idx == 1 {
			break
		}

		heap[idx], heap[idx/2] = heap[idx/2], heap[idx]
		idx /= 2
	}

	return idx
}

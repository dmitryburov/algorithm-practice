package main

/*
ID посылки: 67488184

ПРИНЦИП РАБОТЫ ---
Алгоритм работает по принципу бинарного поиска (за O(log n)) с ипользованием рекурсии.
Поскольку массив частично отсортирован, в базоые случаи и при "делении" нам нужно учитывать "сломанность" стороны.
При определении нужной стороны мы задаем промежуток для рекурсии далее
Добавил комментарии дополнительно в коде.

ВРЕМЕННАЯ СЛОЖНОСТЬ --
Алгоритм работает за О(log n), тк мы проходим только по нужной разделенной части
массива, пока не найдем К или не закончится его "деление"

ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Массив сожержащий n элементов занимает O(n) памяти
Мы работаем в рамках памяти для выделенного массива из n элементов

*/

// brokenSearch основная функция
func brokenSearch(arr []int, k int) int {
	return binarySearch(arr, k, 0, len(arr)-1)
}

// binarySearch бинарный поиск по массиву
func binarySearch(arr []int, k, left, right int) int {

	mid := (left + right + 1) / 2

	// базовые случаи
	if arr[mid] == k {
		return mid
	}
	if left == right {
		return -1
	}

	// рекурсивные случаи
	lf, rg := left, mid-1
	for {
		// если слева ок
		if arr[0] < arr[mid] {
			if k >= arr[0] {
				if k > arr[mid] {
					lf, rg = mid, right
					break
				}
				break
			}

			lf, rg = mid, right
			break
		}

		// если сломан слева
		if k < arr[0] {
			if k < arr[mid] {
				break
			}
			lf, rg = mid, right
			break
		}

		// один проход
		break
	}

	return binarySearch(arr, k, lf, rg)
}

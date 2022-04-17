package main

import (
	"fmt"
	"reflect"
)

func merge(arr []int, lf int, mid int, rg int) []int {
	n1 := mid - lf + 1
	n2 := rg - mid

	l := make([]int, n1)
	r := make([]int, n2)

	for i := 0; i < n1; i++ {
		l[i] = arr[lf+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = arr[mid+1+j]
	}
	i := 0
	j := 0
	k := lf
	for i < n1 && j < n2 {
		if l[i] < r[j] {
			arr[k] = l[i]
			i++
			k++
		} else {
			arr[k] = r[j]
			j++
			k++
		}
	}
	for i == n1 && j < n2 {
		arr[k] = r[j]
		j++
		k++
	}
	for j == n2 && i < n1 {
		arr[k] = l[i]
		i++
		k++
	}

	return arr
}

func merge_sort(arr []int, lf int, rg int) {
	if lf >= rg {
		return
	}

	mid := (rg + lf) / 2

	merge_sort(arr, lf, mid)
	merge_sort(arr, mid+1, rg)

	merge(arr, lf, mid, rg)
}

func main() {
	//a := []int{1, 4, 9, 2, 10, 11}
	//b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}

	//if !reflect.DeepEqual(b, expected) {
	//	panic("WA. Merge")
	//}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}

	fmt.Println(c)
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}

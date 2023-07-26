package main

import "fmt"

func main() {
	fmt.Println(insert_sort([]int{5, 2, 4, 6, 1, 3}))
}

func insert_sort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := 0
		for i = j - 1; i >= 0 && A[i] < key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
	return A
}

аг
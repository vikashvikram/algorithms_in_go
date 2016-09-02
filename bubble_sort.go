package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

func main() {
	a := []int{0, 7, 2, 5, 9, 1, 6, 3, 8}
	b := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	d := []int{3, 4, 3, 1, 1, 1, 4, 4, 4}
	// a := []int{5, 4, 3, 2, 1}
	fmt.Printf("Before Insertion Sorting: %v\n", a)
	sort(a)
	fmt.Printf("After Insertion Sorting: %v\n", a)
	fmt.Printf("Before Insertion Sorting: %v\n", b)
	sort(b)
	fmt.Printf("After Insertion Sorting: %v\n", b)
	fmt.Printf("Before Insertion Sorting: %v\n", c)
	sort(c)
	fmt.Printf("After Insertion Sorting: %v\n", c)
	fmt.Printf("Before Insertion Sorting: %v\n", d)
	sort(d)
	fmt.Printf("After Insertion Sorting: %v\n", d)
}

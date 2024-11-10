package main

import "fmt"

func merge(a []int, p, q, r int) {
	nl := q - p + 1
	nr := r - q

	left := make([]int, nl)
	right := make([]int, nr)

	for i := 0; i < nl; i++ {
		left[i] = a[p+i]
	}
	for i := 0; i < nr; i++ {
		right[i] = a[q+i+1]
	}

	var i, j int
	k := p
	for i < nl && j < nr {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}

	for ; i < nl; i++ {
		a[k] = left[i]
		k++
	}
	for ; j < nr; j++ {
		a[k] = right[j]
		k++
	}
}

func mergeSort(a []int, p, r int) {
	if p >= r {
		return
	}
	q := (r + p) / 2
	mergeSort(a, p, q)
	mergeSort(a, q+1, r)
	merge(a, p, q, r)
}

func main() {
	fmt.Printf("Enter the number of elements: ")
	var n int
	fmt.Scanf("%d", &n)

	fmt.Printf("Enter the elements: ")

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	mergeSort(a, 0, n-1)
	fmt.Println("Sorted array: ", a)
}

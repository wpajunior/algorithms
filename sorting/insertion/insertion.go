package main

import "fmt"


func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >=0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	fmt.Println("Insertion Sort Algorithm")
	fmt.Print("Enter the number of elements: ")
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println("Enter the elements")
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	insertionSort(arr)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}

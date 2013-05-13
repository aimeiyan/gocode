package main

import "fmt"

func exchange(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Quicksort(arr []int) {
	old := make([]int, len(arr))
	copy(old, arr)
	count := len(arr)
	if count <= 1 {
		return
	}

	x := arr[count-1]
	// fmt.Println("Before:", arr, "Count:", count, "X:", x)

	i := 0 //  the first index of the larger part
	for j := 0; j < count-1; j++ {
		if arr[j] <= x {
			exchange(arr, j, i)
			i += 1
		}
		// fmt.Println(arr)
	}

	exchange(arr, i, count-1)

	// fmt.Println(i)

	// tmp := arr[i+1]
	// arr[i+1] = x
	// arr[count-1] = tmp

	// fmt.Println("After", i, arr)
	fmt.Println(old, "---------------", arr[:i], x, arr[i+1:])

	Quicksort(arr[:i])
	Quicksort(arr[i+1:])
}

func maxHeapify(arr []int, i int) {
	size := len(arr)
	left := i*2 + 1
	right := i*2 + 2

	largest := i

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		exchange(arr, i, largest)
		maxHeapify(arr, largest)
	}
}

func HeapSort(arr []int) {
	size := len(arr)
	for i := size / 2; i >= 0; i-- {
		maxHeapify(arr, i)
	}
	// fmt.Println("max heap", arr)

	for i := size - 1; i > 0; i-- {
		exchange(arr, 0, i)
		// fmt.Println(i, arr)
		maxHeapify(arr[:i], 0)
	}
}

func main() {
	arr := []int{11, 17, 6, 3, 9, 2, 8}
	HeapSort(arr)

	fmt.Println("Headsort", arr)

	arr = []int{11, 17, 6, 3, 9, 2, 8}

	// fmt.Println(arr[1:2])
	// fmt.Println(arr)
	Quicksort(arr)
	fmt.Println(arr)
}

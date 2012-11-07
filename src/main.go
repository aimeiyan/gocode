package main

import (
	. "btree"
	"fmt"
)

func main() {
	var tree *Btree = New(4)
	fmt.Println(tree)

	// arr := []int{1, 2, 3, 4}

	// fmt.Println(arr[:2])

	for i := 1; i < 100; i++ {
		tree.Insert(Key(i))

		fmt.Printf("%d------------------------------------\n", i)
		fmt.Println(tree)
		fmt.Printf("%d------------------------------------\n", i)
		fmt.Println()
	}

	fmt.Println(tree)
}

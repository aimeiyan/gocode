package btree_test

import (
	. "btree"
	"fmt"
	"testing"
)

func TestAaaa(t *testing.T) {
	var tree *Btree = New(4)
	fmt.Println(tree)

	for i := 0; i < 100; {
		tree.Insert(Key(i))
	}

	fmt.Printf("%v", *tree)
}

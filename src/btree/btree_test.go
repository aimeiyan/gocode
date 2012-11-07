package btree_test

import (
	. "btree"
	"fmt"
	"testing"
)

func TestAaaa(t *testing.T) {
	var tree *Btree = New(4)
	fmt.Printf("%v", *tree)
}

package btree_test

import (
	. "btree"
	"fmt"
	"testing"
)

func TestAaaa(t *testing.T) {
	var tree *Btree = New(2)

	for i := 0; i < 100; i++ {
		tree.Insert(Key(i))
	}

	// fmt.Println(tree)
}

func TestSplit(t *testing.T) {
	full := &Node{
		Leaf:     true,
		N:        3,
		Keys:     []Key{4, 7, 8},
		Children: make([]*Node, 4),
	}

	p := &Node{
		Leaf:     false,
		N:        1,
		Keys:     []Key{3, 0, 0},
		Children: []*Node{nil, full, nil, nil},
	}

	p.Split(2, 1)
	if full.N != 1 {
		t.Error("full is length should change from 3 to 1")
	}

	if p.N != 2 || p.Keys[1] != 7 || p.Keys[0] != 3 {
		fmt.Println(p)
		t.Error("parent is not correct")
	}

	s := p.Children[2]
	if s == nil || s.N != 1 || s.Keys[0] != 8 {
		t.Error("new child is not correct")
	}
}

func TestInsertRoot(t *testing.T) {
	tree := New(2)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(2)

	r := tree.Root
	if r.N != 3 {
		t.Error("should only have 3 element")
	}

	if r.Keys[0] != 1 || r.Keys[1] != 2 || r.Keys[2] != 10 {
		t.Error("not inserted elements")
	}
}

func TestInsertSplit(t *testing.T) {
	tree := New(2)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(2)

	tree.Insert(5)

	c := tree.Root.Children[1]
	if c == nil || c.N != 2 || c.Keys[0] != 5 {
		t.Error("not correct")
	}
}

func TestInsertSplit2(t *testing.T) {
	tree := New(2)

	for _, v := range []Key{2, 1, 10, 5, 4, 3, 7, 6, 11, 12, 13} {
		tree.Insert(v)
	}

	if tree.Root.N != 1 {
		t.Error("root should have 1 elements")
	}

	c := tree.Root.Children[1]
	if c == nil || c.N != 2 || c.Keys[0] != 7 {
		t.Error("not correct")
	}
}

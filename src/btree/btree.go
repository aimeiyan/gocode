package btree

import (
	// "bytes"
	"fmt"
)

type Key int

type Node struct {
	Leaf     bool
	N        int
	Keys     []Key
	Children []*Node
}

func (x *Node) Search(k Key) (n *Node, idx int) {
	i := 0
	for i < x.N && x.Keys[i] < k {
		i += 1
	}
	if i < x.N && k == x.Keys[i] {
		n, idx = x, i
	} else if x.Leaf == false {
		n, idx = x.Children[i].Search(k)
	}
	return
}

func newNode(n, branch int, leaf bool) *Node {
	return &Node{
		Leaf:     leaf,
		N:        n,
		Keys:     make([]Key, branch*2-1),
		Children: make([]*Node, branch*2),
	}
}

func (x *Node) Split(branch, idx int) { //  idx is Children's index
	y := x.Children[idx] //  x is parent, y is full

	// make a new node, copy y's right most to it
	z := newNode(branch-1, branch, y.Leaf)
	for i := 0; i < branch-1; i++ {
		z.Keys[i] = y.Keys[i+branch]
		z.Children[i] = y.Children[i+branch]
	}
	y.N = branch - 1

	// shift x, add the key children
	for i := x.N; i > idx; i-- {
		x.Children[i] = x.Children[i-1]
		x.Keys[i+1] = x.Keys[i]
	}
	x.Keys[idx] = y.Keys[branch]
	x.Children[idx+1] = z
	x.N += 1
}

func (tree *Btree) Insert(k Key) {
	root := tree.root
	if root.N == 2*tree.branch-1 {
		s := newNode(0, tree.branch, false)
		tree.root = s
		s.Children[0] = root
		s.Split(tree.branch, 0)
		s.InsertNonFull(tree.branch, k)
	} else {
		root.InsertNonFull(tree.branch, k)
	}
}

func (x *Node) InsertNonFull(branch int, k Key) {
	i := x.N
	if x.Leaf {
		for i > 0 && k < x.Keys[i-1] {
			x.Keys[i] = x.Keys[i-1]
			i -= 1
		}
		x.Keys[i] = k
		x.N += 1
	} else {
		for i > 0 && k < x.Keys[i-1] {
			i -= 1
		}
		c := x.Children[i]
		// fmt.Println("key:", k, "idx:", i, "me:", x, x.Keys[i-1], x.Keys[i], "child:", c)
		if c.N == 2*branch-1 {
			x.Split(branch, i)
			if k > x.Keys[i] {
				i += 1
			}
		}
		x.Children[i].InsertNonFull(branch, k)
		// fmt.Println("---insert", c)
	}
}

func space(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

func (x *Node) String() string {
	c := 0

	for c < x.N && x.Children[c] != nil {
		c += 1
	}

	return fmt.Sprintf("{n=%d, Leaf=%v, Keys=%v, Children=%v}",
		x.N, x.Leaf, x.Keys[:x.N], x.Children[:c])
}

func (tree *Btree) String() string {
	return tree.root.String()
}

type Btree struct {
	root   *Node
	branch int
}

func New(branch int) *Btree {
	return &Btree{root: newNode(0, branch, true), branch: branch}
}

func (tree *Btree) Search(k Key) (n *Node, idx int) {
	return tree.root.Search(k)
}

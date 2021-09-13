package main

import (
	"bytes"
	"sort"
)

type treeNode struct {
	name          []byte
	total         uint64
	self          uint64
	childrenNodes []*treeNode
}

func newNode(label []byte) *treeNode {
	return &treeNode{
		name:          label,
		childrenNodes: []*treeNode{},
	}
}

var (
	semicolon = byte(';')
)

type Tree struct {
	root *treeNode
}

func NewTree() *Tree {
	return &Tree{
		root: newNode([]byte{}),
	}
}

func prependBytes(s [][]byte, x []byte) [][]byte {
	s = append(s, nil)
	copy(s[1:], s)
	s[0] = x
	return s
}

func (n *treeNode) insert(targetLabel []byte) *treeNode {
	i := sort.Search(len(n.childrenNodes), func(i int) bool {
		return bytes.Compare(n.childrenNodes[i].name, targetLabel) >= 0
	})

	if i > len(n.childrenNodes)-1 || !bytes.Equal(n.childrenNodes[i].name, targetLabel) {
		child := newNode(targetLabel)
		n.childrenNodes = append(n.childrenNodes, child)
		copy(n.childrenNodes[i+1:], n.childrenNodes[i:])
		n.childrenNodes[i] = child
	}
	return n.childrenNodes[i]
}

func (t *Tree) Insert(key []byte, value uint64) {
	// TODO: maybe we can optimize this slightly by just iterating over it instead of using Split?
	labels := bytes.Split(key, []byte{semicolon})
	node := t.root
	for _, l := range labels {
		buf := make([]byte, len(l))
		copy(buf, l)
		l = buf

		n := node.insert(l)

		node.total += value
		node = n
	}
	node.self += value
	node.total += value
}

func (t *Tree) Iterate(cb func(key []byte, val uint64)) {
	nodes := []*treeNode{t.root}
	prefixes := make([][]byte, 1)
	prefixes[0] = make([]byte, 0)
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		prefix := prefixes[0]
		prefixes = prefixes[1:]

		label := append(prefix, semicolon)
		l := node.name
		label = append(label, l...)

		cb(label, node.self)

		nodes = append(node.childrenNodes, nodes...)
		for i := 0; i < len(node.childrenNodes); i++ {
			prefixes = prependBytes(prefixes, label)
		}
	}
}

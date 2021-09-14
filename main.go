package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"runtime"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(23061912))
}

func benchmarkAllocations(cb func()) {
	mBefore := runtime.MemStats{}
	runtime.ReadMemStats(&mBefore)
	cb()
	mAfter := runtime.MemStats{}
	runtime.ReadMemStats(&mAfter)

	fmt.Printf("  * objects allocated: %d\n", mAfter.HeapObjects-mBefore.HeapObjects)
}

func generateTree(treeWidth, treeDepth, symLen int) *Tree {
	t := NewTree()
	symBuf := make([]byte, symLen)

	for w := 0; w < treeWidth; w++ {
		symbol := []byte("root")
		for d := 0; d < treeDepth; d++ {
			random.Read(symBuf)
			symbol = append(symbol, byte(';'))
			symbol = append(symbol, []byte(hex.EncodeToString(symBuf))...)
			t.Insert(symbol, uint64(random.Intn(100)))
		}

		t.Insert(symbol, uint64(random.Intn(100)))
	}
	return t
}

func main() {
	fmt.Printf("1. allocates one simple tree\n")
	benchmarkAllocations(func() {
		t := NewTree()
		t.Insert([]byte("foo;bar;baz"), 100)
	})

	// we store all trees in this slice so that they don't get garbage collected
	trees := []*Tree{}

	fmt.Printf("2. measures GC time as a function of number of objects on heap\n")
	for i := 10; i < 100; i++ {
		benchmarkAllocations(func() {
			width := i * 20
			depth := i * 50
			fmt.Printf("allocating 100 trees with width %d and depth %d\n", width, depth)
			for j := 0; j < 100; j++ {
				newTree := generateTree(width, depth, 40)
				trees = append(trees, newTree)
			}
		})
	}
}

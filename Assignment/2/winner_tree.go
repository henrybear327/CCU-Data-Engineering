/*
Confirm that all strings are read -> write to file and see if it's the same as the input file
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1-based!

type TreeNode struct {
	value  string
	origin int
}

type WinnerTreeData struct {
	sz   int
	data []TreeNode
	fd   []*os.File
	sc   []*bufio.Scanner
}

func winnerTreeLeftChild(index int) int {
	return index * 2
}

func winnerTreeRightChild(index int) int {
	return index*2 + 1
}

func (w *WinnerTreeData) winnerTreeInternalNodeUpdate(i int) {
	if w.data[winnerTreeLeftChild(i)].origin == -1 {
		w.data[i].origin = w.data[winnerTreeRightChild(i)].origin
		w.data[i].value = w.data[winnerTreeRightChild(i)].value
	} else if w.data[winnerTreeRightChild(i)].origin == -1 {
		w.data[i].origin = w.data[winnerTreeLeftChild(i)].origin
		w.data[i].value = w.data[winnerTreeLeftChild(i)].value
	} else {
		if w.data[winnerTreeLeftChild(i)].value <= w.data[winnerTreeRightChild(i)].value {
			w.data[i].origin = w.data[winnerTreeLeftChild(i)].origin
			w.data[i].value = w.data[winnerTreeLeftChild(i)].value
		} else {
			w.data[i].origin = w.data[winnerTreeRightChild(i)].origin
			w.data[i].value = w.data[winnerTreeRightChild(i)].value
		}
	}
}

func (w *WinnerTreeData) winnerTreePrint() {
	fmt.Printf("\n=====================================\n")
	for i := 0; i < w.sz; i++ {
		fmt.Printf("%v: %v %v\n", i, w.data[i].origin, w.data[i].value)
	}
	fmt.Printf("=====================================\n\n")
}

func (w *WinnerTreeData) winnerTreeInit() {
	total := *config.totalChunks

	// Underlying array size calculation
	// Find the first 2^n which is >= i (leaf level)
	// And then multiply by 2 (internal node levels)
	for i := 0; int(1<<uint(i)) <= total; i++ {
		w.sz = int(1 << uint(i))
	}
	w.sz <<= 2
	fmt.Printf("Winner tree size %v (%v)\n", w.sz, total)

	// 1-based
	// Prepare fd and Scanner
	// TODO: do we need to call make()?
	w.fd = make([]*os.File, 0)
	w.sc = make([]*bufio.Scanner, 0)
	for i := 0; i < w.sz/2; i++ {
		if i < *config.totalChunks {
			w.fd = append(w.fd, openTempFile(i))
			w.sc = append(w.sc, bufio.NewScanner(w.fd[i]))
		} else {
			w.fd = append(w.fd, nil)
			w.sc = append(w.sc, nil)
		}
	}

	// build tree
	w.data = make([]TreeNode, w.sz)
	for i := w.sz - 1; i >= 0; i-- {
		if i >= w.sz/2 { // leaves
			// fmt.Printf("%v %v %v\n", i, i-w.sz/2, w.sc[i-w.sz/2])
			if w.sc[i-w.sz/2] == nil || w.sc[i-w.sz/2].Scan() == false { // empty chunk
				w.data[i].origin = -1
				w.data[i].value = ""
			} else {
				w.data[i].origin = i - w.sz/2
				w.data[i].value = w.sc[i-w.sz/2].Text()
			}
			// fmt.Printf("%v %v: %v value %v\n", i, i-w.sz/2, w.data[i].origin, w.data[i].value)
		} else { // internal nodes
			w.winnerTreeInternalNodeUpdate(i)
		}
	}

	w.winnerTreePrint()
}

func (w *WinnerTreeData) winnerTreeUpdate() {
	if w.winnerTreeIsEmpty() == false {
		index := w.data[1].origin
		// fmt.Printf("index %v\n", index)

		if w.sc[index].Scan() == false {
			w.data[index+w.sz/2].origin = -1
			w.data[index+w.sz/2].value = ""
		} else {
			w.data[index+w.sz/2].value = w.sc[index].Text()
		}

		index += w.sz / 2
		index /= 2
		// fmt.Printf("new index %v\n", index)

		for ; index >= 1; index /= 2 {
			w.winnerTreeInternalNodeUpdate(index)
		}
	}
}

func (w *WinnerTreeData) winnerTreeSize() int {
	return len(w.data)
}

func (w *WinnerTreeData) winnerTreeIsEmpty() bool {
	if w.data[1].origin == -1 {
		return true
	}
	return false
}

func (w *WinnerTreeData) winnerTreeTop() string {
	if w.winnerTreeIsEmpty() == false {
		return w.data[1].value
	}

	panic("Topping a empty winner tree")
}

func (w *WinnerTreeData) winnerTreePop() {
	if w.winnerTreeIsEmpty() == false {
		w.winnerTreeUpdate()
		return
	}

	panic("Popping a empty winner tree")
}

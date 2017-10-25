package main

import (
	"fmt"
)

type TreeNode struct {
	value  string
	origin int
}

type WinnerTreeData struct {
	sz   int
	data []TreeNode
}

var winnerTreeData WinnerTreeData

func initWinnerTree() {
	total := *config.totalChunks
	for i := 0; int(1<<uint(i)) <= total; i++ {
		winnerTreeData.sz = int(1 << uint(i+1))
	}

	fmt.Printf("Winner tree size %v (%v)\n", winnerTreeData.sz, total)

	winnerTreeData.data = make([]TreeNode, winnerTreeData.sz)
}

func updateWinnerTree() {

}

func popWinnerTree() {

	updateWinnerTree()
}

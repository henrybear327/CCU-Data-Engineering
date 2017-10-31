package main

import (
	"sort"
	"sync"
)

type Node struct {
	leftBound, rightBound int
	needSorting           bool
}

var nodeData []Node
var sortedData [][]string

func sortIt(idx int, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Printf("Started node %v sorting\n", idx)
	sort.Strings(sortedData[idx])
	// fmt.Printf("Ended node %v sorting\n", idx)
}

func parallelSort(data []string) {
	// fmt.Printf("Parallel on %v\n", len(data))

	totalNodes := (1 << uint(*config.depth+1))
	nodeData = make([]Node, totalNodes)
	// fmt.Printf("Total %v\n", totalNodes)

	// first pass, get [l, r)
	mergeSort(0, 1, 0, 0, len(data), data)
	// for i := 1; i < totalNodes; i++ {
	// 	fmt.Printf("Node %v: %v %v\n", i, nodeData[i].leftBound, nodeData[i].rightBound)
	// }

	// sort
	var wg sync.WaitGroup
	sortedData = make([][]string, totalNodes)
	for i := 1; i < totalNodes; i++ {
		if nodeData[i].needSorting == false {
			continue
		}

		sortedData[i] = make([]string, nodeData[i].rightBound-nodeData[i].leftBound)
		copy(sortedData[i], data[nodeData[i].leftBound:nodeData[i].rightBound])
		wg.Add(1)
		go sortIt(i, &wg)
	}
	wg.Wait()

	// for i := 1; i < totalNodes; i++ {
	// 	if nodeData[i].needSorting == true {
	// 		for j := 0; j < len(sortedData[i]); j++ {
	// 			fmt.Printf("%v, %v: %v\n", i, j, sortedData[i][j])
	// 		}
	// 	}
	// }

	// merge
	mergeSort(1, 1, 0, 0, len(data), data)

	// for i := 0; i < len(data); i++ {
	// 	fmt.Printf("%v: %v\n", i, data[i])
	// }
}

func mergeSort(passNumber, node, depth, left, right int, data []string) {
	// fmt.Printf("Entering Node %v: %v %v\n", node, left, right)
	nodeData[node].leftBound = left
	nodeData[node].rightBound = right
	if depth == *config.depth {
		if passNumber == 0 {
			nodeData[node].needSorting = true
		} else if passNumber == 1 {
			for i := nodeData[node].leftBound; i < nodeData[node].rightBound; i++ {
				data[i] = sortedData[node][i-nodeData[node].leftBound]
			}
		}
		return
	}

	mid := left + (right-left)/2
	mergeSort(passNumber, node*2, depth+1, left, mid, data)
	mergeSort(passNumber, node*2+1, depth+1, mid, right, data)

	// fmt.Printf("left %v mid %v right %v\n", left, mid, right)

	if passNumber == 1 {
		// fmt.Println("passNumber is 1")
		i := nodeData[node*2].leftBound
		j := nodeData[node*2+1].leftBound
		// fmt.Printf("Bound %v %v, %v %v\n", nodeData[node*2].leftBound, nodeData[node*2].rightBound,
		// 	nodeData[node*2+1].leftBound, nodeData[node*2+1].rightBound)

		tmp := make([]string, nodeData[node*2+1].rightBound-nodeData[node*2].leftBound)

		idx := 0
		for i < nodeData[node*2].rightBound && j < nodeData[node*2+1].rightBound {
			// fmt.Printf("Before %v %v %v\n", i, j, idx)
			if data[i] < data[j] {
				tmp[idx] = data[i]
				i++
			} else {
				tmp[idx] = data[j]
				j++
			}
			idx++
			// fmt.Printf("After %v %v %v\n", i, j, idx)
		}

		for i < nodeData[node*2].rightBound {
			tmp[idx] = data[i]
			i++
			idx++
		}

		for j < nodeData[node*2+1].rightBound {
			tmp[idx] = data[j]
			j++
			idx++
		}

		for i := 0; i < idx; i++ {
			data[nodeData[node*2].leftBound+i] = tmp[i]
			// fmt.Printf("check %v: %v\n", nodeData[node*2].leftBound+i, data[nodeData[node*2].leftBound+i])
		}
	} else {
		// fmt.Println("passNumber is 0")
	}
}

package main

import (
	"sort"
)

var tmpDataForMerge []string

func parallelSortNormal(data []string) {
	// fmt.Printf("Data len = %v\n", len(data))

	tmpDataForMerge = make([]string, len(data))

	done := make(chan bool, 1)
	mergeSortNormal(0, 0, len(data), data, done)
	<-done
}

func mergeSortNormal(depth, left, right int, data []string, done chan bool) {
	// fmt.Printf("Entering Node %v: %v %v\n", node, left, right)
	if depth == *config.depth {
		sort.Strings(data[left:right])
		done <- true
		return
	}

	mid := left + (right-left)/2
	doneLeft := make(chan bool, 1)
	doneRight := make(chan bool, 1)
	go mergeSortNormal(depth+1, left, mid, data, doneLeft)
	go mergeSortNormal(depth+1, mid, right, data, doneRight)
	<-doneLeft
	<-doneRight

	idx := left
	i := left
	j := mid
	for i < mid && j < right {
		// fmt.Printf("Before %v %v %v\n", i, j, idx)
		if data[i] < data[j] {
			tmpDataForMerge[idx] = data[i]
			i++
		} else {
			tmpDataForMerge[idx] = data[j]
			j++
		}
		idx++
		// fmt.Printf("After %v %v %v\n", i, j, idx)
	}

	for i < mid {
		tmpDataForMerge[idx] = data[i]
		i++
		idx++
	}

	for j < right {
		tmpDataForMerge[idx] = data[j]
		j++
		idx++
	}

	for i := left; i < right; i++ {
		data[i] = tmpDataForMerge[i]
		// fmt.Printf("check %v: %v\n", nodeData[node*2].leftBound+i, data[nodeData[node*2].leftBound+i])
	}

	done <- true
}

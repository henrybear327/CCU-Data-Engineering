package main

import (
	"sort"
	"sync"
)

func parallelSortNormal(data []string) {
	var newWG sync.WaitGroup
	newWG.Add(1)
	mergeSortNormal(0, 0, len(data), data, &newWG)
	newWG.Wait()
}

func mySort(in chan []string, out chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	// start := time.Now()
	sort.Strings(<-in)
	// fmt.Println(time.Since(start))
}

func mergeSortNormal(depth, left, right int, data []string, wg *sync.WaitGroup) {
	// fmt.Printf("Entering Node %v: %v %v\n", node, left, right)
	if depth == *config.depth {
		go func(left, right int, wg *sync.WaitGroup) {
			var newData []string
			newChannelIn := make(chan []string, 1)
			newChannelOut := make(chan []string, 1)
			newData = make([]string, len(data[left:right]))
			copy(newData, data[left:right])
			go mySort(newChannelIn, newChannelOut, wg)
			newChannelIn <- newData
			out := <-newChannelOut
			copy(data[left:right], out)
		}(left, right, wg)

		return
	}

	defer wg.Done()
	var newWG sync.WaitGroup

	mid := left + (right-left)/2
	newWG.Add(1)
	go mergeSortNormal(depth+1, left, mid, data, &newWG)

	newWG.Add(1)
	go mergeSortNormal(depth+1, mid, right, data, &newWG)

	newWG.Wait()

	// fmt.Printf("left %v mid %v right %v\n", left, mid, right)

	tmp := make([]string, right-left)

	idx := 0
	i := left
	j := mid
	for i < mid && j < right {
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

	for i < mid {
		tmp[idx] = data[i]
		i++
		idx++
	}

	for j < right {
		tmp[idx] = data[j]
		j++
		idx++
	}

	for i := 0; i < idx; i++ {
		data[left+i] = tmp[i]
		// fmt.Printf("check %v: %v\n", nodeData[node*2].leftBound+i, data[nodeData[node*2].leftBound+i])
	}
}

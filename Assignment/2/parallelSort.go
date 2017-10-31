package main

import (
	"runtime"
	"sort"
	"sync"
)

func mySort(in []string, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Strings(in)
}

var sortedResult []string
var leftPoint []int
var rightPoint []int

func parallelSort(data []string) {
	var wg sync.WaitGroup

	// sort

	// TODO: depth check
	runs := (1 << uint(*config.depth))
	sz := len(data) / runs
	leftPoint = make([]int, runs)
	rightPoint = make([]int, runs)

	newArray := make([][]string, runs)
	for i := 0; i < runs; i++ {
		leftPoint[i] = sz * i
		rightPoint[i] = sz * (i + 1)
		if i == runs-1 {
			rightPoint[i] = len(data)
		}

		newArray[i] = make([]string, len(data[leftPoint[i]:rightPoint[i]]))
		copy(newArray[i], data[leftPoint[i]:rightPoint[i]])
		wg.Add(1)
		go mySort(newArray[i], &wg)
	}
	wg.Wait()

	// merge
	sortedResult = make([]string, len(data))
	idx := int64(0)
	for i := 0; i < runs; i++ {
		for j := 0; j < len(newArray[i]); j++ {
			sortedResult[idx] = newArray[i][j]
			idx++
		}
	}

	wg.Add(1)
	mergesort(0, 1, &wg)
}

func mergesort(dep int, node int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer runtime.UnlockOSThread()

	runtime.LockOSThread()

	var localWG sync.WaitGroup

	if dep >= *config.depth {
		// when threshold is met
		// call system sort

		// sort.Slice(data, func(i, j int) bool {
		// 	return data[i] > data[j]
		// })
		return
	}

	localWG.Add(1)
	go mergesort(dep+1, node*2, &localWG)
	localWG.Add(1)
	go mergesort(dep+1, node*2+1, &localWG)
	localWG.Wait()

	// fmt.Printf("%v %v %v\n", node, dep, node*2-(1<<uint(dep+1)))
	i := leftPoint[node*2-(1<<uint(dep+1))]
	mid := rightPoint[node*2-(1<<uint(dep+1))]
	j := leftPoint[node*2+1-(1<<uint(dep+1))]
	right := rightPoint[node*2+1-(1<<uint(dep+1))]
	idx := 0
	tmp := make([]string, right-i)
	for i < mid && j < right {
		if sortedResult[i] < sortedResult[j] {
			tmp[idx] = sortedResult[i]
			i++
			idx++
		} else {
			tmp[idx] = sortedResult[j]
			j++
			idx++
		}
	}

	for i < mid {
		tmp[idx] = sortedResult[i]
		i++
		idx++
	}

	for j < right {
		tmp[idx] = sortedResult[j]
		j++
		idx++
	}

	for i := 0; i < idx; i++ {
		sortedResult[leftPoint[node*2-(1<<uint(dep+1))]+i] = tmp[i]
	}
}

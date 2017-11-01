package main

import (
	"sort"
	"sync"
)

var tmpDataForMerge []string

func parallelSortNormal(data []string) {
	// fmt.Printf("Data len = %v\n", len(data))

	tmpDataForMerge = make([]string, len(data))

	// var newWG sync.WaitGroup
	// newWG.Add(1)
	done := make(chan bool, 1)
	mergeSortNormal(0, 0, len(data), data, done)
	<-done
	// newWG.Wait()
}

func mySort(in chan []string, out chan []string, wg *sync.WaitGroup) {
	// start := time.Now()
	data := <-in
	sort.Strings(data)
	// for i := 0; i < len(data)-1; i++ {
	// 	if data[i+1] < data[i] {
	// 		panic("sorting data error in mySort")
	// 	}
	// }
	out <- data
	// fmt.Println(time.Since(start))
}

func mergeSortNormal(depth, left, right int, data []string, done chan bool) {
	// fmt.Printf("Entering Node %v: %v %v\n", node, left, right)
	if depth == *config.depth {
		// go func(left, right int, wg *sync.WaitGroup) {
		// 	// fmt.Printf("left %v right %v\n", left, right)
		// 	newChannelIn := make(chan []string, 1)
		// 	newChannelOut := make(chan []string, 1)
		// 	// var newData []string
		// 	// newData = make([]string, len(data[left:right]))
		// 	// copy(newData, data[left:right])
		// 	go mySort(newChannelIn, newChannelOut, wg)
		// 	// newChannelIn <- newData
		// 	newChannelIn <- data[left:right]
		// 	out := <-newChannelOut

		// 	// for i := 0; i < len(out)-1; i++ {
		// 	// 	if out[i+1] < out[i] {
		// 	// 		panic("out sorting data error in mergesort")
		// 	// 	}
		// 	// }

		// 	// if len(data[left:right]) != len(out) {
		// 	// 	panic("data error")
		// 	// }

		// 	// copy(data[left:right], out) // can't use it like this??!!
		// 	for i := 0; i < len(data[left:right]); i++ {
		// 		data[left+i] = out[i]
		// 	}

		// 	// fmt.Println("==== sorted ====")
		// 	// for i := 0; i < len(data[left:right]); i++ {
		// 	// 	fmt.Println(data[left+i])
		// 	// }
		// 	// fmt.Println("====")
		// 	defer wg.Done()
		// }(left, right, wg)

		sort.Strings(data[left:right])
		// defer wg.Done()
		done <- true
		return
	}

	// defer wg.Done()
	// var newWG sync.WaitGroup

	mid := left + (right-left)/2
	// newWG.Add(2)
	doneLeft := make(chan bool, 1)
	doneRight := make(chan bool, 1)
	go mergeSortNormal(depth+1, left, mid, data, doneLeft)
	go mergeSortNormal(depth+1, mid, right, data, doneRight)
	// newWG.Wait()
	<-doneLeft
	<-doneRight

	// fmt.Printf("left %v mid %v right %v\n", left, mid, right)

	// do it once
	// tmp := make([]string, right-left)

	// fmt.Printf("==== before %v %v ====\n", left, right)
	// for i := 0; i < len(data[left:right]); i++ {
	// 	fmt.Println(data[left+i])
	// }
	// fmt.Println("====")

	idx := 0
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

	for i := 0; i < idx; i++ {
		data[left+i] = tmpDataForMerge[i]
		// fmt.Printf("check %v: %v\n", nodeData[node*2].leftBound+i, data[nodeData[node*2].leftBound+i])
	}
	// fmt.Println("==== after ====")
	// for i := 0; i < len(data[left:right]); i++ {
	// 	fmt.Println(data[left+i])
	// }
	// fmt.Println("====")
	done <- true
}

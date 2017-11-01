// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Torture test for goroutines.
// Make a lot of goroutines, threaded together, and tear them down cleanly.

package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup

var total time.Duration

func mySort(in chan []int, out chan []int) {
	defer wg.Done()
	start := time.Now()
	data := <-in
	sort.Ints(data)
	fmt.Println(time.Since(start))
	total += time.Since(start)
	out <- data
}

func mySort2(in []int) {
	defer wg.Done()
	start := time.Now()
	sort.Ints(in)
	fmt.Println(time.Since(start))
	total += time.Since(start)
}

func main() {
	data := make([]int, 0)
	nn := flag.Int("n", 200000000, "Data size")
	pp := flag.Int("p", 8, "Partitions")
	flag.Parse()

	p := *pp
	n := *nn
	for i := n - 1; i >= 0; i-- {
		data = append(data, i)
	}

	fmt.Printf("Partitions %v\n", p)
	wg.Add(p)
	for i := 0; i < p; i++ {
		// sort 1
		// fmt.Printf("i = %v\n", i)
		go func(idx int) {
			// fmt.Printf("func i = %v\n", idx)
			// fmt.Println(n, p, len(data))
			var newData []int
			newChannelIn := make(chan []int, 1)
			newChannelOut := make(chan []int, 1)
			newData = make([]int, len(data[n/p*idx:n/p*(idx+1)]))
			copy(newData, data[n/p*idx:n/p*(idx+1)])
			go mySort(newChannelIn, newChannelOut)
			newChannelIn <- newData
			out := <-newChannelOut
			copy(data[n/p*idx:n/p*(idx+1)], out)
		}(i)

		// sort 2
		// var newData []int
		// newData = make([]int, len(data[n/p*i:n/p*(i+1)]))
		// copy(newData, data[n/p*i:n/p*(i+1)])
		// go mySort2(newData)
	}

	wg.Wait()
	fmt.Println(total)
}

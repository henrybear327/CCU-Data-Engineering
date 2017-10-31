// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Torture test for goroutines.
// Make a lot of goroutines, threaded together, and tear them down cleanly.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func f(left, right chan int) {
	left <- <-right
}

func solve() {
	var n = 10000
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			print("bad arg\n")
			os.Exit(1)
		}
	}

	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	<-leftmost
}

func mySort(in chan []int) {
	data := <-in
	sort.Ints(data)
}

func mySort2(in []int) {
	defer wg.Done()
	sort.Ints(in)
}

func myTest() {
	defer wg.Done()

	rand.Seed(42)
	res := int64(0)
	for i := int64(0); i < 10000000000; i++ {
		res += i
	}
	fmt.Println(res)
}

func main() {
	// solve()

	data := make([]int, 0)
	n := 100000000
	p := 10
	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	wg.Add(p)
	for i := 0; i < p; i++ {
		var ch chan []int
		var newData []int
		newData = make([]int, len(data[n/p*i:n/p*(i+1)]))
		copy(newData, data[n/p*i:n/p*(i+1)])
		go mySort(ch)
		ch <- newData

		// var newData []int
		// newData = make([]int, len(data[n/p*i:n/p*(i+1)]))
		// copy(newData, data[n/p*i:n/p*(i+1)])
		// go mySort2(newData)

		// go myTest()
	}

	wg.Wait()
}

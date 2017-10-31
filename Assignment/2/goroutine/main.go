// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Torture test for goroutines.
// Make a lot of goroutines, threaded together, and tear them down cleanly.

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

func mySort(in chan []int, out chan []int) {
	data := <-in
	sort.Ints(data)
	out <- data
}

func mySort2(in []int) []int {
	sort.Ints(in)
	return in
}

func main() {
	// solve()

	data := make([]int, 0)
	n := 10000000
	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	for i := 0; i < 10; i++ {
		// ch := make(chan []int, 0)
		// out := make(chan []int, 0)
		// go mySort(ch, out)
		// ch <- data[n/10*i : n/10*(i+1)]

		var newData []int
		newData = make([]int, len(data[n/10*i:n/10*(i+1)]))
		copy(newData, data[n/10*i:n/10*(i+1)])

		out := mySort2(newData)
		if len(out) > 0 {
			fmt.Println("Yes")
		}
	}
}

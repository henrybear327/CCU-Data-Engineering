package main

import (
	"runtime"
	"sort"
)

func parallelSort(data []string) {
	ch := make(chan []string)

	in := make(chan []string)
	go mergesort(in, ch, 0)
	in <- data

	data = <-ch
}

func mergesort(in chan []string, out chan []string, dep int) {
	data := <-in
	if dep >= *config.depth {
		// when threshold is met
		// call system sort
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
		out <- data
		return
	}

	runtime.LockOSThread()

	N := len(data)
	res1 := make(chan []string, 1)
	res2 := make(chan []string, 1)

	split1 := make(chan []string)
	split2 := make(chan []string)
	go mergesort(split1, res1, dep+1)
	go mergesort(split2, res2, dep+1)
	split1 <- data[:N/2]
	split2 <- data[N/2:]

	l, r := <-res1, <-res2
	i, j := 0, 0
	for ix := range data {
		switch {
		case i == len(l):
			data[ix] = r[j]
			j++
		case j == len(r):
			data[ix] = l[i]
			i++
		case l[i] > r[j]:
			data[ix] = l[i]
			i++
		default:
			data[ix] = r[j]
			j++
		}
	}

	out <- data

	runtime.UnlockOSThread()
}

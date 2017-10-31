package main

import (
	"runtime"
	"sort"
)

func parallelSort(data []string) {
	out := make(chan []string)

	in := make(chan []string)
	newCopy := make([]string, len(data))
	copy(newCopy, data)
	go mergesort(in, out, 0)
	in <- newCopy

	data = <-out
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
	split1Data := make([]string, len(data[:N/2]))
	split2Data := make([]string, len(data[N/2:]))
	copy(split1Data, data[:N/2])
	copy(split2Data, data[N/2:])
	go mergesort(split1, res1, dep+1)
	go mergesort(split2, res2, dep+1)
	split1 <- split1Data
	split2 <- split2Data

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

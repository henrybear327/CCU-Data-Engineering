package main

import (
	"runtime"
)

func parallelSort(data []string) {
	out := make(chan []string)

	newCopy := make([]string, len(data))
	copy(newCopy, data)
	go mergesort(newCopy, out, 0)

	data = <-out
}

func mergesort(data []string, out chan []string, dep int) {
	if dep >= *config.depth {
		// when threshold is met
		// call system sort

		// sort.Slice(data, func(i, j int) bool {
		// 	return data[i] > data[j]
		// })
		out <- data
		return
	}

	runtime.LockOSThread()

	N := len(data)

	res1 := make(chan []string, 1)
	res2 := make(chan []string, 1)

	// split1 := make(chan []string)
	// split2 := make(chan []string)
	split1Data := make([]string, len(data[:N/2]))
	split2Data := make([]string, len(data[N/2:]))
	copy(split1Data, data[:N/2])
	copy(split2Data, data[N/2:])
	go mergesort(split1Data, res1, dep+1)
	go mergesort(split2Data, res2, dep+1)
	// split1 <- split1Data
	// split2 <- split2Data

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

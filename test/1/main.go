package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}

func createResultFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return file
}

var buffer []string

type Pair struct {
	str   string
	count int
}

func read(scanner *bufio.Scanner, printer *bufio.Writer) {
	buffer := make([]string, 0)
	cnt := make(map[string]int)

	for scanner.Scan() {
		str := scanner.Text()
		strings.Replace(str, "\t", " ", -1)

		buffer = append(buffer, str)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for i := 0; i < len(buffer); i++ {
		splitStr := strings.Split(buffer[i], " ")
		for i := 0; i < len(splitStr); i++ {
			if splitStr[i] == "" {
				continue
			}

			if val, ok := cnt[splitStr[i]]; ok {
				val++
				cnt[splitStr[i]] = val
			} else {
				cnt[splitStr[i]] = 1
			}
		}
	}

	result := make([]Pair, 0)
	for k, v := range cnt {
		// fmt.Printf("key[%s] value[%s]\n", k, v)
		result = append(result, Pair{k, v})
	}

	sort.Slice(result, func(i, j int) bool {
		r1, r2 := result[i], result[j]

		if r1.count == r2.count {
			return r1.str < r2.str
		}
		return r1.count > r2.count
	})

	for i := range result {
		ans := fmt.Sprintf("%v %s\n", result[i].count, result[i].str)
		printer.WriteString(ans)
	}
}

func main() {
	// input := openFile("youtube.s2")
	input, _ := os.Open("/dev/stdin")
	scanner := bufio.NewScanner(input)

	// output := openFile("youtube.s2.out")
	output, _ := os.Create("/dev/stdout")
	printer := bufio.NewWriter(output)
	read(scanner, printer)

	printer.Flush()
	input.Close()
	output.Close()
}

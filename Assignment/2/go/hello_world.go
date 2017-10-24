package main

import (
	"fmt"
	"math"
	"os"
)

// Circle struct
type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func openFiles() {
	file, err := os.Open("number.in")
	if err != nil {
		fmt.Println("No such file")
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat.Name())
	fmt.Println(stat.Size())

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	xs := []float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))

	// var c Circle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := &Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)

	fmt.Println(c.area())

	openFiles()
}

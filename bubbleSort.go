package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randArr := mkArr()
	p := &randArr
	for s := 0; s < len(randArr); s++ {
		for i := len(randArr) - 1; i > 0; i-- {
			if randArr[i] < randArr[i-1] {
				cross(p, i, i-1)
			} else if randArr[i] > randArr[i-1] {
				continue
			}
		}
	}
	fmt.Print(randArr)
}

func random(max int) int {
	var num int
	num = rand.Intn(max)
	return num
}

func mkArr() []int {
	var length int
	fmt.Scan(&length)
	rand.Seed(time.Now().UnixNano())
	array := make([]int, length)
	hoge := 0
	for i := range array {
		hoge++
		array[i] = hoge
	}
	for s := len(array) - 1; s > 0; s-- {
		r := random(length)
		pp := array[s]
		array[s] = array[r]
		array[r] = pp
	}
	return array
}

func cross(array *[]int, index1 int, index2 int) {
	num1 := (*array)[index1]
	(*array)[index1] = (*array)[index2]
	(*array)[index2] = num1
}

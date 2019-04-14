package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	arr := rand.Perm(20)
	length := len(arr)
	arr2 := []int{0}

	for i:=0; i<length; i++ {
		arr2 = append(arr2, 0)
	}

	start := 1

	for _, v := range arr {
		arr2[start] = v
		swim(&arr2, start)
		start++
	}

	fmt.Println(arr)
	fmt.Println(arr2)
	length2 := len(arr2)
	for k, _ := range arr2 {
		if k == 0 {
			continue
		}

		//fmt.Println(arr2[1])
		sink(&arr2, length2)
		length2--
	}
	fmt.Println(arr2)
}

func swim(arr *[]int, curr int){
	tmp := *arr
	if curr == 1 {
		return
	}

	index := 0

	for curr >= 2 {

		index = curr/2
		if tmp[index] > tmp[curr] {
			break
		}
		tmp[index], tmp[curr] = tmp[curr], tmp[index]
		curr = curr/2
	}
	arr = &tmp

	return
}

func sink(arr *[]int, length int) {
	tmp := *arr
	index := 1

	if tmp[index] > tmp[length-1] {
		tmp[index], tmp[length-1] = tmp[length-1], tmp[index]
	}

	for 2*index < length-1 && 2*index+1 < length-1 {
		if tmp[index] < tmp[2*index+1] && tmp[2*index] < tmp[2*index+1] {
			tmp[index], tmp[2*index+1] = tmp[2*index+1], tmp[index]
			index = 2*index+1
			continue
		}
		if tmp[index] < tmp[2*index] {
			tmp[index], tmp[2*index] = tmp[2*index], tmp[index]
			index = 2*index
			continue
		}
		break
	}
	arr = &tmp
	return
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	unsort := rand.Perm(20)
	fmt.Println(unsort)
	QuickSort(&unsort)
	//unsort = []int{1,2,3,4,5}
	fmt.Println(unsort)
	for _,v := range unsort {
		fmt.Println(BinarySearch(unsort, v))
	}
	fmt.Println(BinarySearch(unsort, 21))
}

func BinarySearch(sort []int, value int) int {
	length := len(sort)
	low := 0
	high := length - 1

	return binarySearch(sort, value, low, high)
}

func binarySearch(sort []int, value, low, high int) int {
	index := (high-low) >> 1 + low

	if high >= low {
		if sort[index] == value {
			return index
		} else if sort[index] > value {
			return binarySearch(sort, value, low, index-1)
		} else if sort[index] < value {
			return binarySearch(sort, value, index+1, high)
		}
	}

	return -1
}

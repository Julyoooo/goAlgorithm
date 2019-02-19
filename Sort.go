package main

import "fmt"

func main(){
	unSort := []int{4,5,6,3,2,1}

	SelectSort(&unSort)
	fmt.Println(unSort)

	unSort = []int{3,5,4,1,2,6}

	BubbleSort(&unSort)
	fmt.Println(unSort)

	unSort = []int{3,5,4,1,2,6}

	InsertSort(&unSort)
	fmt.Println(unSort)
}

func SelectSort(u *[]int) {
	unSort := *u
	size := len(unSort)
	
	for i:=0; i<size; i++ {
		for j:=i+1; j<size; j++ {
			if unSort[j] < unSort[i] {
				unSort[i], unSort[j] = unSort[j], unSort[i]
			}
		}
	}
}

func BubbleSort(u *[]int) {
	unSort := *u
	size := len(unSort)
	flag := true

	for i:=1; i<=size; i++ {
		flag = false
		for j:=0; j<size-i; j++ {
			if unSort[j] > unSort[j+1] {
				unSort[j], unSort[j+1] = unSort[j+1], unSort[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertSort(u *[]int) {
	unSort := *u
	size := len(unSort)

	for i:=1; i<size; i++ {
		value := unSort[i]
		for j:=i-1; j>=0; j-- {
			if unSort[j] > value {
				unSort[j+1] = unSort[j]
				unSort[j] = value
			}
		}
	}
}
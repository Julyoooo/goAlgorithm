package main

import (
	"fmt"
)

//func main() {
//
//	src := []int{0,1,2,3,4,5,6,7,8,9}
//	dest := make([]int, 10)
//	perm := rand.Perm(10)
//	for i, v := range perm {
//		dest[v] = src[i]
//	}
//
//	fmt.Println(perm)
//	perm = rand.Perm(10)
//	fmt.Println(perm)
//	fmt.Println(dest)
//
//	unSort := []int{4, 5, 6, 3, 2, 1}
//
//	SelectSort(&unSort)
//	fmt.Println(unSort)
//
//	unSort = []int{3, 5, 4, 1, 2, 6}
//
//	BubbleSort(&unSort)
//	fmt.Println(unSort)
//
//	unSort = []int{3, 5, 4, 1, 2, 6}
//
//	InsertSort(&unSort)
//	fmt.Println(unSort)
//
//	unSort = rand.Perm(10)
//	fmt.Println(unSort)
//	sort := MergeSort(&unSort)
//	fmt.Println("MergeSort:", sort)
//
//	unSort = rand.Perm(10)
//	sort = URMergeSort(&unSort)
//	fmt.Println("URMergeSort:", sort)
//
//	unSort = rand.Perm(50)
//	fmt.Println(unSort)
//	QuickSort(&unSort)
//	fmt.Println("QuickSort:", unSort)
//}

func SelectSort(u *[]int) {
	unSort := *u
	size := len(unSort)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
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

	for i := 1; i <= size; i++ {
		flag = false
		for j := 0; j < size-i; j++ {
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

	for i := 1; i < size; i++ {
		value := unSort[i]
		for j := i - 1; j >= 0; j-- {
			if unSort[j] > value {
				unSort[j+1] = unSort[j]
				unSort[j] = value
			}
		}
	}
}

func MergeSort(u *[]int) []int {
	s := *u
	size := len(s)

	return mergeSort(&s, 0, size-1)
}

func mergeSort(u *[]int, s, e int) []int {
	un := *u
	size := e - s + 1
	ret := make([]int, size)

	un1 := un[s : s+size/2]
	un2 := un[s+size/2 : e+1]

	if size <= 2 {
		ret = p(un1, un2)
	} else {
		ret1 := mergeSort(&un, s, s + size/2-1)
		ret2 := mergeSort(&un, s + size/2, e)
		ret = p(ret1, ret2)
	}

	return ret
}

func p(un1, un2 []int) []int {
	s1 := len(un1)
	s2 := len(un2)
	ret := make([]int, s1+s2)

	if s2 == 0 {
		return un1
	}

	if s1 == 0 {
		return un2
	}

	for i, j, k := 0, 0, 0; ; {
		if un1[i] <= un2[j] {
			ret[k] = un1[i]
			i++
			k++
		} else {
			ret[k] = un2[j]
			j++
			k++
		}

		if i == s1 && j < s2 {
			for {
				ret[k] = un2[j]
				j++
				k++
				if j == s2 {
					break
				}
			}
			break
		}

		if j == s2 && i < s1 {
			for {
				ret[k] = un1[i]
				i++
				k++
				if i == s1 {
					break
				}
			}
			break
		}
	}
	return ret
}

func URMergeSort(u *[]int) []int {
	s := *u
	l := len(s)
	step := 1
	fmt.Println("start:", s)

	for {
		start := 0
		for {
			un1 := s[start : l]
			un2 := s[l:]
			if start + step > l {
				un1 = s[start : l]
				un2 = s[l:]
			} else {
				un1 = s[start : start+step]
				if start + step*2 >= l {
					un2 = s[start+step : l]
				} else {
					un2 = s[start+step : start + step*2]
				}
			}

			tmp := p(un1, un2)

			k := start
			for _,v := range tmp {
				s[k] = v
				k = k + 1
			}
			start = start + step*2
			if start >= l {
				break
			}
		}
		fmt.Println("process:", s)

		step = step * 2
		if step > l {
			break
		}
	}

	return s
}

func QuickSort(u *[]int) {
	s := *u
	quickSort(&s, 0, len(s)-1)
}

func quickSort(u *[]int, s, e int) {
	if s>=e {
		return
	}
	mid := q(u, s, e)
	un := *u

	un[s], un[mid] = un[mid], un[s]
	quickSort(&un, s, mid-1)
	quickSort(&un, mid+1, e)
}

func q(u *[]int, i, j int) int {
	s := *u
	mid := s[i]
	l := j - i + 1
	p := 0

	i = i + 1

	for {
		for s[i] < mid {
			i++
			if i >= l {
				break
			}
		}

		for s[j] > mid {
			j--
			if j < 0 {
				break
			}
		}

		if i >= j {
			p = j
			break
		}

		s[i], s[j] = s[j], s[i]
	}

	return p
}
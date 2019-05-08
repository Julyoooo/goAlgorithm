package main

import (
	"fmt"
)

func main() {
	Calc8Queen(5)
}

func Calc8Queen(size int) {
	var queen [][]int

	for i:=0; i<size; i++ {
		tmp := make([]int, size)
		queen = append(queen, tmp)
	}

	calc8Queen(queen, 0)
}

func calc8Queen(q [][]int, r int) {
	length := len(q)
	if r == length {
		for i:=0; i<length; i++ {
			for j:=0; j<length; j++ {
				fmt.Printf("%v", q[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}

	for col := 0; col < length; col++ {
		if isok(q, r, col) {
			q[r][col] = 1
			calc8Queen(q, r+1)
			q[r][col] = 0
		}
	}

	return
}

func isok(q [][]int, r, c int) bool {
	flag := true
	leftup := c - 1
	rightup := c + 1

	for i := r - 1; i >= 0; i-- {
		if q[i][c] == 1 {
			flag = false
			break
		}
		if leftup >= 0 && q[i][leftup] == 1 {
			flag = false
			break
		}
		if rightup <= len(q)-1 && q[i][rightup] == 1 {
			flag = false
			break
		}
		leftup--
		rightup++
	}

	return flag
}

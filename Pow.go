package main

import "fmt"

func main() {
	fmt.Println("a")
	fmt.Println(myPow(1.732, 2))
	fmt.Println(mySqrt(3))
}

func myPow(x float32, n float32) float32 {
	if n == 0 || x == 1{
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		r := myPow(x, -n)
		return 1/r
	}

	half := int(n) / 2

	if int(n) % 2 == 0 {
		return myPow(x, float32(half)) * myPow(x, float32(half))
	}

	return myPow(x, float32(half)) * myPow(x, float32(half)) * x
}

func mySqrt(x float32) float32 {
	count := 1
	next := float32(0)
	next = (x - next)/2 + next
	cur := myPow(next, 2)

	for {
		if cur > x {
			next = next - (x - next)/2
		}

		if int(cur * 100) == int(x * 100) {
			return next
		}

		if cur < x {
			next = (x - next)/2 + next
		}

		cur = myPow(next, 2)
		fmt.Println("cur:", int(cur * 1000))

		if count > 1000 {
			break
		}

		count++
	}

	return next
}

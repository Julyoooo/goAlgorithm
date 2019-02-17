package main

import (
	"fmt"
	"reflect"
)

func main() {
	var aList [10]int

	var void interface{}

	void = &aList

	for i:=0; i < 10; i++{
		aList[i] = i
	}
	bList := aList[1:]
	bList = append(bList, 10)
	bList = append(bList, 11)
	bList[0] = 100

	l := LinkList{1, nil}

	fmt.Println(l.next)

	fmt.Println(aList)
	fmt.Println(void)
	fmt.Printf("Size of var (reflect.TypeOf.Size): %d\n", reflect.TypeOf(aList[0]).Size())
}

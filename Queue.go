package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type node struct {
	value interface{}
	next *node
}
type Queue struct {
	Head *node
	Size int
	Tail *node
}

const MAX  = 10

func (q *Queue) Init() {
	q.Head = &node{0, nil}
	q.Tail = q.Head
	q.Size = 0
}

func (q *Queue) Enqueue(n node) error{
	if q.Size > MAX {
		return errors.New("Max!")
	}
	if q.Size == 0 {
		q.Head = &n
	}
	q.Tail.next = &n
	q.Tail = q.Tail.next
	q.Size += 1
	return nil
}

func (q *Queue) Dequeue() interface{} {
	if q.Size == 0 {
		fmt.Println("队列为空")
		return nil
	}
	ret := q.Head.value
	value, ok := ret.(int)

	if ok {
		q.Size -= 1
		q.Head = q.Head.next
		return value
	}
	return nil
}

type QueueByArray struct {
	Values []interface{}
	Size int
	Head int
	Tail int
}

func (q *QueueByArray) Init(size int) {
	q.Values = make([]interface{}, size)
	q.Size = 0
	q.Head = 0
	q.Tail = 0
}

func (q *QueueByArray) Enqueue(value interface{}) {
	if q.Size == 0 {
		q.Values[q.Head] = value
	}
	q.Values[q.Tail] = value
	q.Tail += 1
}

// todo: 数组队列搬移

func main() {
	q := Queue{}

	q.Init()

	q.Enqueue(node{1, nil})

	q.Enqueue(node{2, nil})

	q.Enqueue(node{3, nil})

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue() == nil)

	fmt.Println("Queue:↑")

	fmt.Println("QueueByArray:")
	qa := QueueByArray{}

	qa.Init(2)
	qa.Enqueue(1)
	qa.Enqueue(2)
	fmt.Println(qa.Values, qa.Head, qa.Tail)
}
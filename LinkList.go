package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	prev *Node
	value int
	next *Node
}

func (head *Node) Append(n *Node) {
	head.next = n
	n.prev = head
}

func (head *Node) Insert(p *Node, n *Node) error{
	cur := head.next
	if cur == nil {
		return errors.New("empty list")
	}
	for {
		if cur == p {
			n.next = cur.next
			n.prev = cur
			cur.next.prev = n
			cur.next = n
			break
		}
		cur = cur.next
	}

	return nil
}

func (head *Node) String() string{
	var ret string

	ret = "head"
	cur := head.next
	for {
		if cur == nil{
			break
		}
		ret = ret + "->" + strconv.Itoa(cur.value)
		cur = cur.next
	}
	return ret
}

func (head *Node) Reverse() { // 1->2->3
	cur := head.next
	t := Node{nil, 0, nil}

	for {
		if cur == nil {
			break
		}
		n := cur.next
		cur.next = t.next
		t.next = cur
		cur = n
	}
	head.next = t.next
}

func main()  {
	var head Node
	var l Node
	var n Node
	var insert Node

	head = Node{nil, 0, nil}
	l = Node{nil, 1, nil}
	n = Node{nil,2 , nil}
	insert = Node{nil,3, nil}

	head2 := Node{nil, 0, nil}
	a := Node{nil, 1, nil}
	b := Node{nil, 2, nil}
	c := Node{nil, 3, nil}
	//d := Node{nil, 3, nil}
	e := Node{nil, 2, nil}
	f := Node{nil, 1, nil}


	head.Append(&l)
	l.Append(&n)
	head.Insert(&l, &insert)

	head2.Append(&a)
	a.Append(&b)
	b.Append(&c)
	//c.Append(&d)
	c.Append(&e)
	e.Append(&f)
	fmt.Println(head2.String())

	p1 := head2.next
	p2 := head2.next.next

	var p Node
	for {

		if p2 == nil { // odd
			fmt.Println("odd", p1.value)
			p.next = p1
			break
		}

		if p2.next == nil { // even
			fmt.Println("even", p1.value)
			p.next = p1.next
			break
		}
		p1 = p1.next
		p2 = p2.next.next
	}
	fmt.Println(p.String())
	fmt.Println(head2.String())
	p.Reverse()

	fmt.Println(p.String())
	fmt.Println(head2.String())

	tmp1 := head2.next
	tmp2 := p.next

	for {
		if tmp1.value != tmp2.value {
			fmt.Println("不是回文", tmp1.value, tmp2.value)
			break
		}

		if tmp1.next == nil || tmp2.next == nil {
			p.Reverse()
			fmt.Println("是回文", head2.String())
			break
		}

		tmp1 = tmp1.next
		tmp2 = tmp2.next
	}
}

func (head *Node) isPalindromic() {
	p1 := head.next
	p2 := head.next.next

	var p Node
	for {

		if p2 == nil { // odd
			p.next = p1
			break
		}

		if p2.next == nil { // even
			p.next = p1.next
			break
		}
		p1 = p1.next
		p2 = p2.next.next
	}

	p.Reverse()

	tmp1 := head.next
	tmp2 := p.next

	for {
		if tmp1.value != tmp2.value {
			fmt.Println("不是回文", tmp1.value, tmp2.value)
			break
		}

		if tmp1.next == nil || tmp2.next == nil {
			p.Reverse()
			fmt.Println("是回文", head.String())
			break
		}

		tmp1 = tmp1.next
		tmp2 = tmp2.next
	}
}
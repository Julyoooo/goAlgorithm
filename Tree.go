package main

import (
	"fmt"
)

type TreeNode struct {
	value interface{}
	left *TreeNode
	right *TreeNode
}

func (t *TreeNode) String() string{
	var ret string

	ret = "here " + t.value.(string)

	if t.left != nil {
		ret += t.left.value.(string)
	}
	return ret
}

func (t *TreeNode) Travers() {
	if t == nil {
		return
	}
	fmt.Println(t.value)
	t.left.Travers()
	t.right.Travers()
}

func BFS(t *TreeNode) {
	if t == nil {
		return
	}
	q := QueueByArray{}
	q.Init(10)
	q.Enqueue(*t)

	for q.Size > 0 {
		c := q.Dequeue().(TreeNode)
		fmt.Println(c.value)
		if c.left != nil {
			q.Enqueue(*c.left)
		}
		if c.right != nil {
			q.Enqueue(*c.right)
		}
	}
}

func DFS(t *TreeNode) {
	if t == nil {
		return
	}

	root := *t

	fmt.Println(root.value)
	if root.left != nil {
		DFS(root.left)
	}
	if root.right != nil {
		DFS(root.right)
	}

	return
}

func main() {
	root := TreeNode{1, nil, nil}
	t1 := TreeNode{2, nil, nil}
	root.left = &t1
	t2 := TreeNode{3, nil, nil}
	root.right = &t2
	t3 := TreeNode{4, nil, nil}
	t1.left = &t3
	t4 := TreeNode{5, nil, nil}
	t1.right = &t4
	t5 := TreeNode{6, nil, nil}
	t2.left = &t5
	t6 := TreeNode{7, nil, nil}
	t2.right = &t6
	fmt.Println("*******1*******")
	fmt.Println("*****/***\\*****")
	fmt.Println("****2*****3****")
	fmt.Println("***/*\\***/*\\***")
	fmt.Println("**4***5*6***7**")

	//fmt.Println("-------前中后序遍历--------")
	//root.Travers()
	fmt.Println("-------广度优先（队列实现）--------")
	BFS(&root)
	fmt.Println("-------深度优先（栈实现）--------")
	DFS(&root)

	m := map[string]int{}
	m["a"] = 10
	m["b"] = 21

	for k,v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(m["a"])
}

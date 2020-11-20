package main

import "fmt"

func main() {
	queue := new(Queue)
	queue.Add(12)
	queue.Add(45)
	queue.Add(19)
	queue.Add(49)
	queue.Add(34)
	queue.Print()
	queue.Remove()
	queue.Print()

}

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyError")
		return 0
	}
	return q.head.value
}

func (q *Queue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (q *Queue) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *Queue) Remove() {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyError")
		return
	}
	q.head = q.head.next
	q.size--

}

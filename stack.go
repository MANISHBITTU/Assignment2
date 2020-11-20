package main

import "fmt"

func main() {
	stk := new(Stack)
	stk.Push(20)
	stk.Push(35)
	stk.Push(22)
	stk.Push(33)
	stk.Print()
	stk.Pop()
	stk.Print()
}

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return 0
	}
	return s.head.value
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return
	}
	s.head = s.head.next
	s.size--
}

func (s *Stack) Print() {
	temp := s.head
	fmt.Print("Value stored in stck are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

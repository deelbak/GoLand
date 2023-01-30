package LIFO

import "fmt"
type Node struct {
	Val int
	next *Node
	prev *Node
}

type Stack struct {
	head *Node
	tail *Node
}
func (s *Stack) Push(node *Node){
	if s.head == nil{
		s.head = node
	 	s.head.next = nil
		s.tail = s.head
	}
}
func (s *Stack) Peek(node *Node) int{
	return s.tail.Val
}
func (s *Stack) Pop() {
	s.tail = s.tail.prev
	s.tail.next = nil
}
func (s *Stack) Clear(){
	s.head = nil
}
func (s *Stack) Contains(val int) bool{
	cur := s.head
	for cur != nil{
		if (cur.Val == val){
			return true
		}
		cur = cur.next
	}
	return false
}
func (s *Stack) Increment(val int) {
	cur := s.head
	for cur != nil{
		cur.Val++
		cur = cur.next
	}
	
}
func (s *Stack) Print(){
	cur := s.head
	for cur != nil{
		fmt.Println(cur.Val)
		cur = cur.next
	}
}
func (s *Stack) PrintReverse(){
	cur := s.tail
	for cur !=nil{
		fmt.Println(cur.Val)
		cur = cur.prev
	}
}
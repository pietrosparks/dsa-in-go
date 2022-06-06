package main

import "fmt"

type arrstack struct {
	items []int
	top   int
}

func (s *arrstack) push(val int) {
	s.top++
	s.items = append(s.items, val)
}

func (s *arrstack) pop() {
	if s.isEmpty() {
		return
	}

	s.top--
}

func (s *arrstack) getTop() int {
	fmt.Printf("\n This is the top: %d", s.top)
	return s.top
}

func (s *arrstack) isEmpty() bool {
	return s.top < 0
}

func (s *arrstack) print() {
	for i := 0; i < s.top; i++ {
		fmt.Printf("\n %#+v", s.items[i])
	}
}

type node struct {
	data int
	next *node
}

type llstack struct {
	head *node
	len  int
}

func (l *llstack) push(val int) {
	n := new(node)
	n.data = val
	n.next = l.head

	l.head = n
	l.len++
}

func (l *llstack) pop() {
	if l.isEmpty() {
		return
	}

	head := l.head.next
	l.head = head
	l.len--
}

func (l *llstack) getTop() int {
	fmt.Printf("\n This is the top: %d", l.len)
	return l.len
}

func (l *llstack) isEmpty() bool {
	return l.len == 0
}

func (l *llstack) print() {
	current := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("\n %#+v", current.data)
		current = current.next
	}
}

func main() {
	s := new(llstack)

	s.push(10)
	s.push(20)
	s.push(45)
	s.push(90)
	s.push(89)
	s.pop()
	s.pop()
	s.push(900)
	s.getTop()
	s.print()

	as := new(llstack)

	as.push(10)
	as.push(20)
	as.push(45)
	as.push(90)
	as.push(89)
	as.pop()
	as.pop()
	as.push(900)
	as.getTop()
	as.print()
}

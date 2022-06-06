package main

import (
	"fmt"
	"strings"
)

type stack interface {
	push(val string)
	pop() string
	peek() string
	isEmpty() bool
}

// Array Implementation
type arrstack struct {
	items []string
	top   int
}

func newArrStack() *arrstack {
	return &arrstack{
		items: []string{},
		top:   -1,
	}
}

func (s *arrstack) push(val string) {
	s.items = append(s.items, val)
	s.top++
}

func (s *arrstack) pop() string {
	top := s.peek()

	if s.isEmpty() {
		return top
	}

	s.top--
	return top
}

func (s *arrstack) peek() string {
	return s.items[s.top]
}

func (s *arrstack) isEmpty() bool {
	return s.top < 0
}

func (s *arrstack) print() {
	for i := 0; i < s.top; i++ {
		fmt.Printf("\n %#+v", s.items[i])
	}
}

// Linked List Implementation
type node struct {
	data string
	next *node
}

type llstack struct {
	head *node
	len  int
}

func (l *llstack) push(val string) {
	n := new(node)
	n.data = val
	n.next = l.head

	l.head = n
	l.len++
}

func (l *llstack) pop() string {
	top := l.peek()

	if l.isEmpty() {
		return top
	}

	head := l.head.next
	l.head = head
	l.len--

	return top
}

func (l *llstack) peek() string {
	return l.head.data
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
	ll := newReverser(new(llstack)).reverse("Weirdo")
	arr := newReverser(newArrStack()).reverse("Weirdo")

	fmt.Printf("\nLinkedList: %s", ll)
	fmt.Printf("\nArray: %s", arr)

}

// Reverse a string using stack
type reverser struct {
	stack
}

func newReverser(r stack) reverser {
	return reverser{r}
}

func (r reverser) reverse(word string) string {
	split := strings.Split(word, "")

	for _, s := range split {
		r.push(s)
	}

	reversed := make([]string, len(split))

	for _, _ = range split {
		reversed = append(reversed, r.pop())
	}

	return strings.Join(reversed, "")
}

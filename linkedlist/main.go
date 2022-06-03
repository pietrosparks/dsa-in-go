package main

// Node is a node attached to a linked list
type node struct {
	data int
	next *node
}

// LinkedList is te=he actual list constituting of patched nodes
type linkedList struct {
	head *node
	len  int
}

func main() {
	l := new(linkedList)
	l.insertBack(8)
	l.insertBack(5)
	l.insertBack(7)
	l.insertBack(9)
	l.insertAt(0, 10)
	l.reverse()
	l.print()
}

// insertBack adds a node to the end of a linked list
func (l *linkedList) insertBack(val int) {
	// create a node to be attached to the linked list
	n := new(node)
	n.data = val

	// if the linked list is empty i.e len = 0. Add to head and exit
	if l.len == 0 {
		l.head = n
		l.len++
		return
	}

	// on first iteration [starter] will be the head of the linked list which calls [next] if not nil and moves to the next linked list till reach end
	starter := l.head
	for i := 0; i < l.len; i++ {

		// if end of linked list attach
		if starter.next == nil {
			starter.next = n
			l.len++
			return
		}

		// if not end of linked list, call next
		starter = starter.next
	}
}

// insertFront adds a node to the front of a linked list
func (l *linkedList) insertFront(val int) { //  -> [data | next ] -> [data | next ]
	n := new(node)
	n.data = val

	// if the linked list is empty i.e len = 0. Add to head and exit
	if l.len == 0 {
		l.head = n
		l.len++
		return
	}

	// set former head as new tail on new node
	// set new node as head
	n.next = l.head
	l.head = n
	l.len++
	return
}

// insertAt adds a node at a specified position
func (l *linkedList) insertAt(pos int, val int) {
	n := new(node)
	n.data = val

	starter := l.head

	for i := 0; i <= pos; i++ {
		if starter.next == nil {
			return
		}

		// if position beginning of node, use insertFront
		if pos == 0 {
			l.insertFront(val)
			return
		}

		// if position end of node or specified beyond, use insertBack
		if pos == l.len-1 || pos > l.len-1 {
			l.insertBack(val)
			return
		}

		// else if we are at the index before the specified position
		// reposition our node's next pointer to point to the position specified's node
		// i.e [0 -> 1 -> 3 -> 10] with pos: 2 tells us to move 3 forward to pos: 3 and insert our val at pos: 2
		// so when we are at val 1 (pos: 1), we set the tail of our new node to point to 3, and the tail of 1 to point to 2
		if i == pos-1 {
			n.next = starter.next // set the tail of our new node to point to 3
			starter.next = n      // the tail of 1 to point to 2
			l.len++
			return
		}

		starter = starter.next
	}
}

func (l *linkedList) deleteAt(pos int) {
	starter := l.head

	for i := 0; i <= pos; i++ {

		if pos > l.len-1 {
			return
		}

		if pos == 0 {
			starter.next = nil
			l.len--
			return
		}

		if i == pos-1 {
			starter.next = starter.next.next
			l.len--
			return
		}

		starter = starter.next
	}
}

func (l *linkedList) pop() {
	starter := l.head
	for i := 0; i < l.len; i++ {

		if i == l.len-2 {
			starter.next = nil
			return
		}

		starter = starter.next
	}
}

func (l *linkedList) print() {
	// attempt to show data from the head
	starter := l.head

	for i := 0; i < l.len; i++ {
		println(starter.data)

		// if next is nil (end of linked list) eject
		if starter.next == nil {
			return
		}

		// else move to next value
		starter = starter.next
	}
}

func (l *linkedList) reverse() {
	var prev, current, next *node
	current = l.head

	for i := 0; i < l.len; i++ {
		// if next is nil (end of linked list) eject
		if current == nil {
			return
		}

		// [0 -> 3 -> 5] prev = nil, current = 0, next = nil
		// we need to set current.next (->) to point to prev (nil)
		// but if we do that, we lose the ability to traverse to 3, so we store current.next in a variable next
		// after pointing 0 -> nil, we need to be able to point 3 to 0. current is 0 and we want to point 3 to 0
		// so we make 0 prev (prev = current), then make current=next (current = 3)
		// so its basically two phases. first phase is the actual reversing. next phase is preparing the (next) to be reversed in next iteration
		// first iteration [0 -> 3 -> 5]
		// 				   p[c -> n]
		// first iteration 	[p<-c n] [nil <- 0 3 -> 5]
		// second iteration [nil <- 0 <- 3 5]
		// 					[nil <- p <- c n]
		// so its a sliding window, where p initially sits outside the window p[c->n]
		// then eventually becomes [p<-c n]

		// store [0 -> 1] next is 1
		next = current.next // 1 // store address of next node
		current.next = prev // 0 // set address of current node to previous node

		prev = current // nil = 0
		current = next // 0 = 1
	}

	// reassign head to be the prev [which is current] so it can be traversed
	l.head = prev

}

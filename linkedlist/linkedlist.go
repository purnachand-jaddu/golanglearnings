package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
}

type IDataStructure interface {
	insert(val int)

	print()

	delete(val int)

	search(val int) bool

	sum() int
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) insert(val int) {
	if l.head == nil {
		l.head = &Node{}
		l.head.value = val
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		newNode := Node{}
		newNode.value = val
		current.next = &newNode
	}

}

func (l *LinkedList) print() {
	node := l.head
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) sum() int {
	sum := 0
	node := l.head
	for node != nil {
		sum += node.value
		node = node.next
	}
	return sum
}

func (l *LinkedList) delete(val int) {
	parent := l.head
	if parent.value == val {
		l.head = parent.next
	} else {
		current := parent.next
		for current != nil && current.value != val {
			current = current.next
			parent = parent.next
		}
		if current == nil {
			return
		} else {
			parent.next = current.next
		}
	}
}

func (l *LinkedList) search(val int) bool {
	node := l.head
	for node != nil {
		if node.value == val {
			return true
		}
		node = node.next
	}
	return false
}

func TestLinkedList() {
	list := LinkedList{}
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.print()
	list.delete(1)
	list.print()
	list.delete(3)
	list.print()
	list.insert(4)
	list.insert(5)
	list.insert(100)
	list.print()
	fmt.Println(list.search(100))
	fmt.Println("Sum is :", list.sum())
}

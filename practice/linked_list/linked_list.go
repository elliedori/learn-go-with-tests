package linkedlist

import "fmt"

type Node struct {
	Body string
	Next *Node
}

type LinkedList struct {
	Length int
	Start  *Node
}

func CreateLinkedList(value string) *LinkedList {
	startingNode := Node{
		Body: value,
		Next: nil,
	}
	l := LinkedList{
		Length: 1,
		Start:  &startingNode,
	}
	return &l
}

func (l *LinkedList) Insert(value string) {
	newTerminalNode := Node{
		Body: value,
		Next: nil,
	}

	currentNode := l.Start

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = &newTerminalNode
	l.Length++
}

func (l *LinkedList) Delete(value string) {

	var previousNode *Node
	currentNode := l.Start

	for currentNode.Body != value {
		if currentNode.Next == nil {
			panic("matching value not found")
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	previousNode.Next = currentNode.Next
	l.Length--

}

func (l *LinkedList) Display() string {
	var visualList string
	currentNode := l.Start
	for currentNode.Next != nil {
		visualList += currentNode.Body
		currentNode = currentNode.Next
	}
	visualList += currentNode.Body
	fmt.Println(visualList)
	return visualList
}

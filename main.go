package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next     *Node
	prev     *Node
	Name     string
	Purchase int
}

type LinkList struct {
	head   *Node
	tail   *Node
	length int
}

func (L LinkList) Len() int {
	return L.length
}

func (L LinkList) Display() {
	if L.length == 0 {
		fmt.Println("The List is empty")
	}
	for L.head != nil {
		fmt.Printf("%v -> ", L.head)
		L.head = L.head.prev
	}
	fmt.Println()
}

func (L *LinkList) AddHead(Node *Node) {
	if L.length == 0 {
		L.head, L.tail = Node, Node
		L.length++
	} else {
		Node.prev = L.head
		L.head.next = Node
		L.head = Node
		L.length++
	}
}

func (L *LinkList) AddTail(Node *Node) {
	if L.length == 0 {
		L.head, L.tail = Node, Node
		L.length++
	} else {
		L.tail.prev = Node
		Node.next = L.tail
		L.tail = Node
		L.length++
	}
}

func (L *LinkList) PopHead() error {
	if L.Len() == 0 {
		return errors.New("The List is already empty - nothing to remove")
	} else if L.Len() == 1 {
		*L = LinkList{}
		return nil
	} else {
		L.head = L.head.prev
		L.head.next.prev = nil
		L.head.next = nil
		L.length--
		return nil
	}
}

func (L *LinkList) PopTail() error {
	if L.Len() == 0 {
		return errors.New("The List is already empty - nothing to remove")
	} else if L.Len() == 1 {
		*L = LinkList{}
		return nil
	} else {
		L.tail = L.tail.next
		L.tail.prev.next = nil
		L.tail.prev = nil
		L.length--
		return nil
	}
}

func main() {

	List := LinkList{}
	Node1 := Node{
		Name:     "Pervyy",
		Purchase: 100,
		// next nil prev 2
	}
	List.AddTail(&Node1)
	Node2 := Node{
		Name:     "Vtoroy",
		Purchase: 200,
		// next 1 prev 3
	}
	List.AddTail(&Node2)
	/*Node3 := Node{
		Name: "Tretyy",
		Purchase: 300,
		// next 2 prev nil
	}
	List.AddTail(&Node3)
	*/

	fmt.Println(List)
	List.PopHead()
	fmt.Println(Node1)
	fmt.Println(List)
	List.Display()
}

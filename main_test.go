package main

import (
	"errors"
	"testing"
)

/*
func TestLinkList_AddHead(t *testing.T) {
	type fields struct {
		head   *Node
		tail   *Node
		length int
	}
	type args struct {
		Node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &LinkList{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
		})
	}
}

func TestLinkList_AddTail(t *testing.T) {
	type fields struct {
		head   *Node
		tail   *Node
		length int
	}
	type args struct {
		Node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &LinkList{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
		})
	}
}
*/

func TestLinkList_PopHead(t *testing.T) {
	NodeTest1 := Node{
		Name:     "Pervyy",
		Purchase: 100,
	}
	NodeTest2 := Node{
		Name:     "Vtoroy",
		Purchase: 200,
	}
	////////////////////// TEST 1 (1 element in list)
	{
		List := LinkList{&NodeTest1, &NodeTest1, 1}
		_ = List.PopHead()
		want := (LinkList{})
		got := List
		if got != want {
			t.Errorf("Test #1 FAILED. Wanted: %v; Got: %v", want, got)
		}
	}
	////////////////////// TEST 2 (2 elements in list)
	{
		//NodeTest1.prev = &NodeTest2
		//NodeTest2.next = &NodeTest1
		List := LinkList{}
		List.AddHead(&NodeTest2)
		List.AddHead(&NodeTest1)
		//List := LinkList{&NodeTest1, &NodeTest2, 2}
		_ = List.PopHead()
		want := LinkList{&NodeTest2, &NodeTest2, 1}
		if List != want || NodeTest2.next != nil || NodeTest1.prev != nil {
			t.Errorf("Test #2 FAILED. Wanted: %v; Got: %v", want, List)
		}
	}
	////////////////////// TEST 3 (0 elements in list)
	{
		List := LinkList{}
		gotError := List.PopHead()
		wantError := errors.New("The List is already empty - nothing to remove")
		if gotError == nil {
			t.Errorf("Test #3 FAILED. Wanted error: %v, Got error: %v", wantError, gotError)
		}
	}
}

func TestLinkList_PopTail(t *testing.T) {
	NodeTest1 := Node{
		Name:     "Pervyy",
		Purchase: 100,
	}
	NodeTest2 := Node{
		Name:     "Vtoroy",
		Purchase: 200,
	}
	////////////////////// TEST 1 (1 element in list)
	{
		List := LinkList{&NodeTest1, &NodeTest1, 1}
		_ = List.PopTail()
		want := (LinkList{})
		if List != want {
			t.Errorf("Test #1 FAILED. Wanted: %v; Got: %v", want, List)
		}
	}
	////////////////////// TEST 2 (2 elements in list)
	{
		NodeTest1.prev = &NodeTest2
		NodeTest2.next = &NodeTest1
		List := LinkList{&NodeTest1, &NodeTest2, 2}
		_ = List.PopTail()
		want := LinkList{&NodeTest1, &NodeTest1, 1}
		if List != want || NodeTest1.prev != nil || NodeTest2.next != nil {
			t.Errorf("Test #2 FAILED. Wanted: %v; Got: %v", want, List)
		}
	}
	////////////////////// TEST 3 (0 elements in list)
	{
		List := LinkList{}
		gotError := List.PopTail()
		wantError := errors.New("The List is already empty - nothing to remove")
		if gotError == nil {
			t.Errorf("Test #3 FAILED. Wanted error: %v, Got error: %v", wantError, gotError)
		}
	}
}

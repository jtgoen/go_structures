package go_structures

import (
	"errors"
)

type LLNode struct {
	value    int
	nextNode *LLNode
}

func (llnode LLNode) Value() int {
	return llnode.value
}

func (llnode LLNode) NextNode() *LLNode {
	return llnode.nextNode
}

func (llnode LLNode) Members() (int, *LLNode) {
	return llnode.value, llnode.nextNode
}

type LinkedList struct {
	head   *LLNode
	tail   *LLNode
	length int
}

func (llist LinkedList) Head() *LLNode {
	return llist.head
}

func (llist LinkedList) Tail() *LLNode {
	return llist.tail
}

func (llist LinkedList) Length() int {
	return llist.length
}

func (llist LinkedList) Members() (*LLNode, *LLNode, int) {
	return llist.head, llist.tail, llist.length
}

func NewLLNode(value int) *LLNode {
	return &LLNode{value, nil}
}

func NewLinkedList(headptr *LLNode) *LinkedList {
	return &LinkedList{headptr, headptr, 1}
}

func (llist LinkedList) At(position int) (*LLNode, error) {

	if position > llist.length-1 || position < 0 {
		return nil, errors.New("Index Out of Range")
	}

	wantedNodeptr := llist.head
	for i := 0; i < position; i++ {
		if wantedNodeptr == llist.tail {
			return nil, errors.New("Reached end of list")
		}
		wantedNodeptr = wantedNodeptr.nextNode
	}
	return wantedNodeptr, nil
}

func (llist *LinkedList) Add(newnodeptr *LLNode) {
	llist.tail.nextNode = newnodeptr
	llist.tail = newnodeptr
	llist.length += 1
}

func (llist *LinkedList) Delete(position int) error {

	wantedNodeptr, e := llist.At(position)
	if e != nil {
		return errors.New("Delete Failed: " + e.Error())
	}

	if wantedNodeptr == llist.head {
		llist.head = wantedNodeptr.nextNode
		wantedNodeptr = nil
		if llist.length > 0 {
			llist.length -= 1
		}
		return nil
	}

	prevNodeptr, e := llist.At(position - 1)
	if e != nil {
		return errors.New("Delete Failed. Error accessing preceding node. " + e.Error())
	}

	if wantedNodeptr == llist.tail {
		llist.tail = prevNodeptr
		wantedNodeptr = nil
		llist.length -= 1
		return nil
	}

	nextNodeptr, e := llist.At(position + 1)
	if e != nil {
		return errors.New("Delete Failed. Error accessing following node. " + e.Error())
	}

	prevNodeptr.nextNode = nextNodeptr
	wantedNodeptr = nil
	llist.length -= 1
	return nil
}

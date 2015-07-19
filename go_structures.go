package go_structures

import (
	"errors"
)

type LLNode struct {
	value    int
	nextNode *LLNode
}

type LinkedList struct {
	head   *LLNode
	tail   *LLNode
	length int
}

func NewLinkedList(headptr *LLNode) *LinkedList {
	return &LinkedList{headptr, headptr, 1}
}

func (llist LinkedList) add(newnodeptr *LLNode) {
	llist.tail.nextNode = newnodeptr
	llist.tail = newnodeptr
}

type llnodeArgError struct {
	problem string
}

func (llist LinkedList) at(position int) (*LLNode, error) {

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

func (llist LinkedList) delete(position int) error {

	wantedNodeptr, e := llist.at(position)
	if e != nil {
		return errors.New("Delete Failed: " + e.Error())
	}

	if position == 0 {
		llist.head = wantedNodeptr.nextNode
		wantedNodeptr = nil
		if llist.length > 0 {
			llist.length -= 1
		}
		return nil
	}

	prevNodeptr, e := llist.at(position - 1)
	if e != nil {
		return errors.New("Delete Failed. Error accessing preceding node. " + e.Error())
	}

	if wantedNodeptr == llist.tail {
		llist.tail = prevNodeptr
		wantedNodeptr = nil
		llist.length -= 1
		return nil
	}

	nextNodeptr, e := llist.at(position + 1)
	if e != nil {
		return errors.New("Delete Failed. Error accessing following node. " + e.Error())
	}

	prevNodeptr.nextNode = nextNodeptr
	wantedNodeptr = nil
	llist.length -= 1
	return nil
}

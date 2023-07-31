package main

import (
	"fmt"
)

type listElem struct {
	name	string
	next	*listElem
}

type linkedList struct {
	size	int
	head	*listElem
}

func initLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) PushFront(name string)() {
	elem := &listElem {
		name: name,
	}

	if 0 == ll.size {
		ll.head = elem
	} else {
		elem.next = ll.head
		ll.head = elem
	}

	ll.size++
	return
}

func (ll *linkedList) PushBack(name string) {
	elem := &listElem {
		name: name,
	}

	if ll.head == nil {
		ll.head = elem
	} else {
		currentElem := ll.head
		for currentElem.next != nil {
			currentElem = currentElem.next
		}
		currentElem.next = elem
	}
	ll.size++
	return
}

func (ll *linkedList) PopFront() error {
	if ll.head == nil {
		return fmt.Errorf("PopFront error: list is empty")
	}
	ll.head = ll.head.next
	ll.size--
	return nil
}

func (ll *linkedList) PopBack() error {
	if ll.head == nil {
		return fmt.Errorf("PopBack error: list is empty")
	}
	var prevElem *listElem
	currentElem := ll.head
	for currentElem.next != nil {
		prevElem = currentElem
		currentElem = currentElem.next
	}
	if prevElem != nil {
		prevElem.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return nil
}

func (ll *linkedList) GetHead() (string, error) {
	if nil == ll.head {
		return "", fmt.Errorf("GetHead error: list is empty")
	}
	return ll.head.name, nil
}

func (ll *linkedList) GetSize() int {
	return ll.size
}

func (ll *linkedList) PrintLinkedList() error {
	if nil == ll.head {
		return fmt.Errorf("print list error: list is empty")
	}
	currentElem := ll.head
	fmt.Print(currentElem.name)
	for currentElem.next != nil {
		currentElem = currentElem.next
		fmt.Print(" ", currentElem.name)
	}
	fmt.Println("")
	return nil
}

func (ll *linkedList) Reverse() {
	if ll.size < 2 {
		return
	}
	currentElem := ll.head
	var prevElem, nextElem *listElem
	for currentElem != nil {
		nextElem = currentElem.next
		currentElem.next = prevElem
		prevElem = currentElem
		currentElem = nextElem
	}
	ll.head = prevElem
}

func main() {
	ll := initLinkedList()
	fmt.Println("the init a new list, list size: ", ll.GetSize())
	err := ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print ("add back a1, printing list: ")
	ll.PushBack("a1")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	fmt.Print ("add back a2, printing list: ")
	ll.PushBack("a2")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PopFront()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopFront success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	fmt.Print ("add front a3, printing list: ")
	ll.PushFront("a3")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PopBack()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopBack success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	fmt.Print ("add front a4, printing list: ")
	ll.PushFront("a4")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\tlist size: ", ll.GetSize())


	fmt.Print ("GetHead: ")
	val, err := ll.GetHead()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	err = ll.PopFront()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopFront success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PopFront()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopFront success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PopFront()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopFront success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PopBack()
	if err != nil {
		fmt.Println(err)
		} else {
		fmt.Print ("PopBack success, printing list: ")
		err = ll.PrintLinkedList()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("\tlist size: ", ll.GetSize())

	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print ("add back 2 elements, printing list: ")
	ll.PushBack("a1")
	ll.PushBack("a2")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	ll.Reverse()
	fmt.Print ("reverse list, printing list: ")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print ("add back 3 elements, printing list: ")
	ll.PushBack("a3")
	ll.PushBack("a4")
	ll.PushBack("a5")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}

	ll.Reverse()
	fmt.Print ("reverse list, printing list: ")
	err = ll.PrintLinkedList()
	if err != nil {
		fmt.Println(err)
	}
}

//go run linkedList.go 
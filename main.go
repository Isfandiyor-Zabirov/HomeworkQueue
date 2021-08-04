package main

import (
	"fmt"
)

type List struct {
	root   *ListNode
	end    *ListNode
	length int
}

type ListNode struct {
	Prev           *ListNode
	Next           *ListNode
	Name           string
	PurchaseAmount int
}

//	Задаем объём очереди
func (receiver *List) len() int  {
	return receiver.length
}

//	Add добавляет нового человека в очереди
func (receiver *List) Add(node ListNode)  {
	if receiver.len() == 0 {
		receiver.root = &node
		receiver.end = &node
		receiver.length++
		return
	}
	LastElement := receiver.end
	receiver.end.Next = &node
	receiver.end = &node
	receiver.end.Prev = LastElement
	receiver.length++
}
//	Remove удаляет человека из списка
func (receiver *List) Remove(node ListNode) {
	if receiver.len() == 1 { 	//если в очереди только 1 чел
		receiver.root = nil
		receiver.end = nil
		receiver.length = 0
		return
	} else { 		//т.е., если в очереди более 1 чел:
		current := receiver.root
		current.Next.Prev = nil
		current.Next = current.Next.Next
		receiver.length--
	}
	return
}

//	PrintList выводит полный список людей в очереди
func (receiver *List) PrintList()  {
	current := receiver.root
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

//	getFirstElement выводит первого человека в очереди
func (receiver *List) getFirstElement(node *ListNode) {
	firstElement := receiver.root
	fmt.Printf("The first one in queue is %v \n", firstElement)
}

// getLastElement выводит последнего человека в очереди
func (receiver *List) getLastElement(node *ListNode)  {
	lastElement := receiver.end
	fmt.Printf("The last one in queue is %v \n", lastElement)
}

func main() {

	person := ListNode{
		Prev:           nil,
		Next:           nil,
		Name:           "Isfandiyor",
		PurchaseAmount: 150,
	}

	queue := List{}
	queue.Add(person)
//	fmt.Println(queue)

	person = ListNode{
		Prev:           nil,
		Next:           nil,
		Name:           "Umed",
		PurchaseAmount: 110,
	}

	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:           nil,
		Next:           nil,
		Name:           "Surush",
		PurchaseAmount: 118,
	}
	queue.Add(person)

queue.getFirstElement(&person)
queue.getLastElement(&person)

	queue.Remove(person)
	queue.PrintList()

}

package main

import (
	"fmt"
	"sync"
)

// 数组实现的队列
type ArrayQueue struct {
	queue []interface{}
	size  int
	mux   sync.Mutex
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		queue: make([]interface{}, 0),
	}
}

func (aq *ArrayQueue) Push(v interface{}) {
	aq.mux.Lock()
	defer aq.mux.Unlock()

	if aq.queue == nil {
		aq.queue = make([]interface{}, 0)
	}
	aq.queue = append(aq.queue, v)

	aq.size += 1
}

func (aq *ArrayQueue) Pop() interface{} {
	aq.mux.Lock()
	defer aq.mux.Unlock()

	if aq.size == 0 {
		return nil
	}

	v := aq.queue[0]
	aq.queue = aq.queue[1:]

	aq.size -= 1

	return v
}

func (aq *ArrayQueue) Size() int {
	return aq.size
}

// 链表实现的队列
type LinkQueue struct {
	start *Node
	end   *Node
	size  int
	mux   sync.Mutex
}

type Node struct {
	next  *Node
	value interface{}
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		start: new(Node),
	}
}

func (lq *LinkQueue) Push(v interface{}) {
	lq.mux.Lock()
	defer lq.mux.Unlock()

	if lq.size == 0 {
		lq.start = new(Node)
		lq.start.value = v
		lq.end = lq.start
		lq.size += 1
		return
	}
	newNode := new(Node)
	newNode.value = v
	lq.end.next = newNode

	lq.end = newNode
	lq.size += 1

	return
}

func (lq *LinkQueue) Pop() interface{} {
	lq.mux.Lock()
	defer lq.mux.Unlock()

	if lq.size == 0 {
		return nil
	}
	newStart := lq.start.next
	v := lq.start.value

	lq.start = newStart
	lq.size -= 1

	return v
}

func (lq *LinkQueue) Size() int {
	return lq.size
}

func main() {

	//queue := NewArrayQueue()
	queue := NewLinkQueue()
	queue.Push("Tom")
	queue.Push("Jack")
	queue.Push("Lisa")
	fmt.Printf("%v\n", queue.Size())
	fmt.Printf("%v\n", queue.Pop())
	fmt.Printf("%v\n", queue.Pop())

	fmt.Printf("%v\n", queue.Size())
	fmt.Printf("%v\n", queue.Pop())
	fmt.Printf("%v\n", queue.Pop())
}

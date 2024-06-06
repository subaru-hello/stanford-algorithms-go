package main

import (
	"container/heap"
	"fmt"
)
type Item struct {
	value int
	priority int
	index int
}

// Heap Data Structure
type PriorityQueue []*Item

// Basic Heap Methods
func (pq *PriorityQueue) Len() int { return len(*pq) }
// to extract the minimum value of the two
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].priority < (*pq)[j].priority }
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	lastItem := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return lastItem
}

// update the priority of an item in the heap
func (pq *PriorityQueue) updatePq(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {

	n := 

	heap1 := make([]int, 0)
	heap2 := make([]int, 0)

	fmt.Println("Hello, World!")
}


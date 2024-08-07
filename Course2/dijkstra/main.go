// Dijkstra's Algorithm
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// create A Graph Structure

type Graph struct {
	nodes map[int][]Edge
}

type Edge struct {
	toNode int
	weight int
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int][]Edge),
	}
}

// an instance method to add an edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.nodes[from] = append(g.nodes[from], Edge{toNode: to, weight: weight})
}

// create A Heap Interface
type Item struct {
    node    int // The value of the item
    priority int // The priority of the item (lower value means higher priority)
	index int // The index of the item in the heap

}

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

// Dikjstra's Algorithm
func Dijkstra(graph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	pq := make(PriorityQueue, len(graph.nodes))
	i := 0
	for node := range graph.nodes {
		if node == start {
			distances[node] = 0
			pq[i] = &Item{node: node, priority: 0, index: i}
			i++
		}else{
			// initialize all other distances to infinity
			distances[node] = math.MaxInt64
			pq[i] = &Item{node: node, priority: math.MaxInt64, index: i}
			i++
		}
	}
	heap.Init(&pq)
	
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		currentDistance := distances[current.node]

		// nodes: [{0, 3}, {1, 1}, {2, 3}]
		// update pq with the min distance among the current node's neighbors
		for _, edge := range graph.nodes[current.node] {
			newDist := currentDistance + edge.weight

			// pick the min distance
			if newDist < distances[edge.toNode] {
				distances[edge.toNode] = newDist
				pq.updatePq(&Item{node: edge.toNode, priority: newDist},newDist)
			}
		}
	}
	return distances
}

// 1. Initialize a priority queue with the source vertex as the first element.
// 2. Initialize a set to keep track of the vertices that have been visited.
// 3. While the queue is not empty, dequeue the vertex with the lowest distance from the source vertex, and add it to the set.
// 4. For each vertex that is adjacent to the dequeued vertex, update its distance if it is shorter than the current distance.
// 5. Repeat until all vertices have been visited.
// 6. The shortest path from the source vertex to each other vertex can be found in the distances array.

func main(){
	graph := NewGraph()
	edges := []struct{
		from, to, weight int
	}{
		{1,2,1},
		{1,3,4},
		{2,3,2},
		{2,4,3},
		{3,4,5},
		{3,5,2},
		{4,5,3},
	}
	for _, edge := range edges {
		graph.AddEdge(edge.from, edge.to, edge.weight)
	}
	startNode := 1
	distances := Dijkstra(graph, startNode)

	fmt.Println("Distances from start Node to all other nodes:", distances)

	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
}

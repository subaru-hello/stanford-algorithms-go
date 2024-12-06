## Karger Min Cut Overview

The Karger min cut algorithm is a well-known graph search algorithm. It finds the minimum number of crossing edges between two partitions of a graph by repeatedly contracting edges until only two vertices remain.

For instance, given a graph G with an initial vertex v and a next vertex w, the function recursively contracts edges between these vertices until the process concludes. Suppose there is a graph comprising 4 vertices and 4 edges connecting these vertices. Given that this is an undirected graph, the minCut function would be applied as minCut(G, [v, w]).

## Running Time

In the worst-case scenario, the Karger min cut algorithm runs in
𝑂(𝑁^2log⁡𝑁) time because the total calculation time reduces by half with each iteration of the computation process. To the contrary, in the best-case scenario, the running time would be O(N^2) due to fewer iterations required when the graph is simplified more quickly with each contraction.

## Pseudocode

1. Read and Parse the Graph: First, we need to read the graph from the file and parse it into an adjacency list representation.
2. Contract Edges: We repeatedly contract randomly chosen edges until only two vertices remain.
3. Repeat the Process: We repeat the contraction process multiple times to ensure a higher probability of finding the minimum cut.

## Simple Algorithm

Here is the code written in Go:

```
package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Graph struct {
    vertices map[int][]int
}

// Contracts the graph by merging the two vertices v1 and v2
func (g *Graph) contract(edge [2]int) {
    v1, v2 := edge[0], edge[1]

    // to contract the graph size
    g.vertices[v1] = append(g.vertices[v1], g.vertices[v2]...)
    for _, v := range g.vertices[v2] {
        edges := g.vertices[v]
        for i := range edges {
            if edges[i] == v2 {
                edges[i] = v1
            }
        }
    }
    delete(g.vertices, v2)
}

func (g *Graph) randomEdge() [2]int {
    v1 := rand.Intn(len(g.vertices)) + 1
    adj := g.vertices[v1]
    v2 := adj[rand.Intn(len(adj))]
    return [2]int{v1, v2}
}

func kargerMinCut(g *Graph) int {
    for len(g.vertices) > 2 {
        edge := g.randomEdge()
        g.contract(edge)
    }
    var edges [][]int
    for _, adj := range g.vertices {
        edges = append(edges, adj)
    }
    return len(edges[0])
}

func main() {
    rand.Seed(time.Now().UnixNano())

    g := &Graph{
        vertices: map[int][]int{
            1: {2, 3, 4},
            2: {1, 3, 4},
            3: {1, 2, 4},
            4: {1, 2, 3},
        },
    }

    minCut := kargerMinCut(g)
    fmt.Println("Minimum Cut:", minCut)
}
```

## Summary

The Karger min cut algorithm is effective for finding the minimum cut in an undirected graph through edge contraction. Despite its simplicity, it achieves good performance with a running time of
O(N^2logN) in the worst case and in the best case. The provided Go implementation demonstrates the basic steps to execute

// find a min cut
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Graph struct {
    vertices map[int][]int
}

/** 
 - merges the adjacency lists of v1 and v2
 - update all other vertices' adjacency lists to reflect the new vertex, self-loop, and edge.
 - removes the edge (v1, v2)
*/
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

/**
 - returns a random edge from the graph
*/
func (g *Graph) randomEdge() [2]int {
	var v1 int

	// find a v1 with non-empty adjacency list
	for v1 = range g.vertices {
		if len(g.vertices[v1]) > 0 {
			break
		}
	}

	// find a random adjacent vertex
	v2 := g.vertices[v1][rand.Intn(len(g.vertices[v1]))]

	return [2]int{v1, v2}
}

/**
 - repeatedly contracts the graph until only 2 vertices are left
 - returns the number of edges in the min cut
 - the minimum number of edges required to remove to disconnect the graph into two components
*/
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
## Overview

The Breadth First Search algorithm is a graph search algorithm that explores all vertices of a graph in breadth-first order.

It is useful when the graph is unweighted and the shortest path to a vertex is needed.
if you could find a right path at first, you can stop the search and return the path.
The time complexity is O(V + E), because each vertex is only visited once and each edge is only considered once. The Worst Case is O(V^2) if the graph is a dense graph.

## Problem Statement

It is a such a basic graph algorithm, so this algorithm has been used in many situations such as finding the shortest path to a vertex, detecting cycles, and finding the connected components of a graph.

## Algorithm

1. Initialize a queue Q with the source vertex s.
2. Initialize a set S to keep track of the vertices that have been processed.
3. While Q is not empty, remove the vertex with the minimum key value from Q, call it u, and add it to S.
4. For each vertex v that is adjacent to u and not in S, update the key value of v to the minimum of the current key value of v and the weight of the edge (u, v).

## Data Structure

Traversing the graph with BFS adopts FIFO strategy, so it is a Queue data structure.

- Queue Q: to keep track of the vertices to be processed.
- Set S: to keep track of the vertices that have been processed.

## Pseudocode

```
function BFS(G, s):
    // Initialize a queue Q with the source vertex s.
    Q <- Queue()
    // Initialize a set S to keep track of the vertices that have been processed. true means that the vertex has been processed.
    S <- Set()
    // Enqueue the source vertex s to Q.
    Q.enqueue(s)
    // While Q is not empty, dequeue a vertex u from Q and add it to S.
    while Q is not empty:
        u <- Q.dequeue()
        S.add(u)
        // For each vertex v that is adjacent to u and not in S, enqueue v to Q.
        for each v in G.adj[u]:
            if v not in S:
                Q.enqueue(v)
```

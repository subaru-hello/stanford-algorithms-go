## Overview

The Depth First Search algorithm is a graph search algorithm that explores all vertices of a graph in depth-first order.

It is useful when the graph is unweighted and the shortest path to a vertex is needed.

DFS applies LIFO approach, so it is a Stack data structure.

## Problem Statement

Here is the sample situations that use DFS.

- Find the shortest path to a vertex in a graph.
- Detect cycles in a graph.
- Find the connected components of a graph.

### Pseudocode

```
function DFS(G, s):
    S <- Set()
    for each vertex v in G:
        if v not in S:
            DFS(G, v)
```

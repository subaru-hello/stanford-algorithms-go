## Overview

input: a directed graph G = (V, E)
output: for each v in V, the shortest path from vertex s to v in G

## Algorithm

1. Initialize a min heap Q with all vertices in V, with the key value being the shortest distance from the source vertex s to each vertex.
2. Initialize a set S to keep track of the vertices that have been processed.
3. While Q is not empty, remove the vertex with the minimum key value from Q, call it u, and add it to S.
4. For each vertex v that is adjacent to u and not in S, update the key value of v to the minimum of the current key value of v and the weight of the edge (u, v).

### Dijkstra's Algorithm

**Dijkstra's Algorithm** is a graph search algorithm that finds the shortest path from a starting node to all other nodes in a weighted graph with non-negative weights. It is widely used in network routing protocols and various applications where the shortest path is needed.

### How It Works

1. **Initialization:**
   - Set the distance to the starting node to 0 and the distance to all other nodes to infinity.
   - Create a priority queue (min-heap) and insert the starting node with a distance of 0.
2. **Processing Nodes:**
   - While the priority queue is not empty:
     - Extract the node with the smallest distance (current node) from the priority queue.
     - For each neighboring node of the current node, calculate the tentative distance from the starting node to this neighbor.
     - If the calculated tentative distance is less than the known distance to the neighbor, update the neighbor's distance and insert the neighbor into the priority queue with the updated distance.
3. **Termination:**
   - The algorithm terminates when all nodes have been processed and their shortest distances from the starting node are known.

### Example

Consider a graph with nodes A, B, C, D, and E with the following edges and weights:

```
A --1-- B
A --4-- C
B --2-- C
B --5-- D
C --1-- D
D --3-- E

```

**Step-by-Step Execution:**

1. **Initialization:**
   - Distances: A=0, B=∞, C=∞, D=∞, E=∞
   - Priority Queue: [(A, 0)]
2. **Processing Nodes:**
   - Extract A (distance 0):
     - Update B: 0 + 1 = 1 (Distances: A=0, B=1, C=∞, D=∞, E=∞; Queue: [(B, 1)])
     - Update C: 0 + 4 = 4 (Distances: A=0, B=1, C=4, D=∞, E=∞; Queue: [(B, 1), (C, 4)])
   - Extract B (distance 1):
     - Update C: 1 + 2 = 3 (Distances: A=0, B=1, C=3, D=∞, E=∞; Queue: [(C, 3), (C, 4)])
     - Update D: 1 + 5 = 6 (Distances: A=0, B=1, C=3, D=6, E=∞; Queue: [(C, 3), (C, 4), (D, 6)])
   - Extract C (distance 3):
     - Update D: 3 + 1 = 4 (Distances: A=0, B=1, C=3, D=4, E=∞; Queue: [(C, 4), (D, 6), (D, 4)])
   - Extract D (distance 4):
     - Update E: 4 + 3 = 7 (Distances: A=0, B=1, C=3, D=4, E=7; Queue: [(D, 6), (E, 7)])
   - Extract remaining nodes, but no shorter paths are found.
3. **Final Distances:**
   - A=0, B=1, C=3, D=4, E=7

The shortest path from A to E is A -> C -> D -> E with a distance of 7.

### Applications

- **Routing protocols:** Used in various network routing protocols like OSPF (Open Shortest Path First) to find the shortest path in a network.
- **Maps and navigation:** Used in GPS systems and mapping applications to find the shortest route between locations.
- **Network design:** Used to optimize the layout and design of networks, ensuring minimal latency and cost.

### Advantages

- **Efficiency:** Can handle large graphs efficiently with appropriate data structures like min-heaps.
- **Simplicity:** Conceptually straightforward and easy to implement.

### Limitations

- **Non-negative weights:** Cannot handle graphs with negative edge weights. For such cases, algorithms like Bellman-Ford are more suitable.

Overall, Dijkstra's Algorithm is a fundamental tool in graph theory and has broad applications in computing and optimization problems.

### code in C++

```cpp
#include <iostream>      // For input/output
#include <vector>        // For using vectors
#include <queue>         // For priority queue
#include <utility>       // For using pairs
#include <limits>        // For defining infinity

using namespace std;

typedef pair<int, int> pii; // Define a pair type (distance, vertex)
const int INF = numeric_limits<int>::max(); // Define infinity as the maximum int value

void dijkstra(const vector<vector<pii>>& graph, int start, vector<int>& distances) {
    int n = graph.size(); // Get the number of vertices in the graph
    distances.assign(n, INF); // Initialize distances with infinity
    distances[start] = 0; // Distance to the starting vertex is 0

    // Priority queue to process vertices by the smallest distance (min-heap)
    priority_queue<pii, vector<pii>, greater<pii>> pq;
    pq.push({0, start}); // Push the starting vertex into the queue

    while (!pq.empty()) {
        int dist = pq.top().first; // Get the smallest distance
        int u = pq.top().second; // Get the corresponding vertex
        pq.pop(); // Remove it from the queue

        // If the distance in the queue is greater than the current known distance, skip
        if (dist > distances[u]) continue;

        // Iterate over neighbors of the current vertex
        for (const auto& edge : graph[u]) {
            int v = edge.first; // Neighbor vertex
            int weight = edge.second; // Weight of the edge

            // Relax the edge if a shorter path is found
            if (distances[u] + weight < distances[v]) {
                distances[v] = distances[u] + weight;
                pq.push({distances[v], v}); // Push the neighbor with updated distance
            }
        }
    }
}

int main() {
    int n, m;
    cout << "Enter the number of vertices and edges: ";
    cin >> n >> m; // Read the number of vertices and edges

    vector<vector<pii>> graph(n); // Create a graph with n vertices
    cout << "Enter the edges (u v weight):\\n";
    for (int i = 0; i < m; ++i) {
        int u, v, weight;
        cin >> u >> v >> weight; // Read the edges
        graph[u].push_back({v, weight}); // Add edge to the graph
        graph[v].push_back({u, weight}); // If undirected graph, add both directions
    }

    int start;
    cout << "Enter the starting vertex: ";
    cin >> start; // Read the starting vertex

    vector<int> distances; // Vector to store distances
    dijkstra(graph, start, distances); // Call Dijkstra's algorithm

    // Output the shortest distances from the starting vertex
    cout << "Shortest distances from vertex " << start << ":\\n";
    for (int i = 0; i < n; ++i) {
        if (distances[i] == INF) {
            cout << "Vertex " << i << ": INF" << endl;
        } else {
            cout << "Vertex " << i << ": " << distances[i] << endl;
        }
    }

    return 0;
}

```

### **Explanation**

1. **Graph Representation:**
   - The graph is represented as an adjacency list where each element is a vector of pairs. Each pair contains a vertex and the weight of the edge connecting to that vertex.
2. **Priority Queue:**
   - A priority queue (min-heap) is used to efficiently get the vertex with the smallest distance at each step.
3. **Initialization:**
   - The distances to all vertices are initialized to infinity, except for the starting vertex, which is set to 0.
4. **Processing:**
   - The main loop extracts the vertex with the smallest distance from the priority queue.
   - For each neighbor of this vertex, it checks if a shorter path can be found and updates the distance accordingly.
   - The neighbor is then pushed into the priority queue with the updated distance.
5. **Input and Output:**
   - The program reads the number of vertices and edges, followed by the edges themselves.
   - It then runs Dijkstra's Algorithm from the specified starting vertex and outputs the shortest distances to all other vertices.

### Techniques Used in the Code

1. **Graph Representation with Adjacency List:**

   - The graph is represented as a vector of vectors of pairs. Each pair represents a connected vertex and the weight of the edge.
   - This is efficient for storing sparse graphs and allows easy iteration over the neighbors of a vertex.

   ```cpp
   typedef pair<int, int> pii;
   vector<vector<pii>> graph(n);

   ```

2. **Priority Queue (Min-Heap):**

   - A priority queue is used to always process the vertex with the smallest known distance.
   - In C++, the `priority_queue` by default is a max-heap, so we use `greater<pii>` to make it a min-heap.

   ```cpp
   priority_queue<pii, vector<pii>, greater<pii>> pq;

   ```

3. **Initialization:**

   - All distances are initialized to infinity (`INF`) to represent that initially, no vertex is reachable.
   - The distance to the starting vertex is set to 0.

   ```cpp
   const int INF = numeric_limits<int>::max();
   distances.assign(n, INF);
   distances[start] = 0;

   ```

4. **Dijkstra's Algorithm:**

   - The main idea of Dijkstra's Algorithm is to repeatedly extract the vertex with the smallest distance and update the distances to its neighbors if shorter paths are found.
   - The priority queue ensures that the vertex with the smallest distance is processed first.

   ```cpp
   while (!pq.empty()) {
       int dist = pq.top().first;
       int u = pq.top().second;
       pq.pop();

       if (dist > distances[u]) continue;

       for (const auto& edge : graph[u]) {
           int v = edge.first;
           int weight = edge.second;

           if (distances[u] + weight < distances[v]) {
               distances[v] = distances[u] + weight;
               pq.push({distances[v], v});
           }
       }
   }

   ```

5. **Efficient Edge Relaxation:**
   - For each neighbor of the current vertex, the algorithm checks if a shorter path to that neighbor exists (`distances[u] + weight < distances[v]`).
   - If a shorter path is found, the distance is updated, and the neighbor is pushed into the priority queue with the new distance.

### Synonyms and Terminology in C++

1. **Graph Representation:**
   - **Vector of vectors of pairs**: `vector<vector<pii>>`
     - Synonym: Adjacency list
   - **Pair of integers**: `pair<int, int>`
     - Synonym: Tuple, ordered pair
2. **Priority Queue:**
   - **Min-Heap**: `priority_queue<pii, vector<pii>, greater<pii>>`
     - Synonym: Min-priority queue, ascending priority queue
3. **Initialization:**
   - **Infinity**: `numeric_limits<int>::max()`
     - Synonym: Maximum integer value, INT_MAX
4. **Algorithm Steps:**
   - **Extract the vertex**: `pq.top()`
     - Synonym: Dequeue the smallest element
   - **Update distances**: `distances[u] + weight < distances[v]`
     - Synonym: Relax edges, update path length
5. **Graph Edges and Weights:**
   - **Edge relaxation**: Updating the shortest path estimate for a vertex
   - **Neighbor**: Adjacent vertex, connected vertex

### Summary

- **Graph Representation:** Adjacency list using vectors of pairs.
- **Priority Queue:** Min-heap for efficiently processing vertices.
- **Initialization:** Distances initialized to infinity.
- **Dijkstra's Algorithm:** Extract min, relax edges, update distances.
- **Efficient Edge Relaxation:** Check and update shortest paths.

This code and explanation should give you a comprehensive understanding of how Dijkstra's Algorithm is implemented in C++.

### Correctness

Dijkstra's Algorithm Implementation

```python
import heapq

def dijkstra(graph, start):
    # Step 1: Initialize distances and priority queue
    distances = {node: float('infinity') for node in graph}
    distances[start] = 0
    priority_queue = [(0, start)]

    while priority_queue:
        current_distance, current_node = heapq.heappop(priority_queue)

        # Nodes can get added to the priority queue multiple times. We only
        # process a node the first time we remove it from the priority queue.
        if current_distance > distances[current_node]:
            continue

        # Step 2: Explore the neighbors of the current node
        for neighbor, weight in graph[current_node].items():
            distance = current_distance + weight

            # Only consider this new path if it's better
            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(priority_queue, (distance, neighbor))

    return distances

# Example graph
graph = {
    'A': {'B': 1, 'C': 4},
    'B': {'A': 1, 'C': 2, 'D': 5},
    'C': {'A': 4, 'B': 2, 'D': 1},
    'D': {'B': 5, 'C': 1}
}

# Running Dijkstra's algorithm from node 'A'
distances_from_a = dijkstra(graph, 'A')
print(distances_from_a)

```

### Test Case to Prove Correctness

We'll use a simple graph for which we can manually verify the shortest paths.

1. **Graph**:
   - 'A' -> 'B' (weight 1)
   - 'A' -> 'C' (weight 4)
   - 'B' -> 'C' (weight 2)
   - 'B' -> 'D' (weight 5)
   - 'C' -> 'D' (weight 1)
2. **Expected Shortest Paths** from 'A':
   - 'A' to 'A': 0
   - 'A' to 'B': 1
   - 'A' to 'C': 3 (A -> B -> C)
   - 'A' to 'D': 4 (A -> B -> C -> D)

The output of the algorithm should match these expected distances.

```python
# Expected distances for validation
expected_distances = {
    'A': 0,
    'B': 1,
    'C': 3,
    'D': 4
}

# Check correctness
assert distances_from_a == expected_distances, "The algorithm's output is incorrect!"
print("Dijkstra's algorithm implementation is correct!")

```

### Output

If the implementation is correct, you should see the printed dictionary with the shortest paths from node 'A', and the assertion should pass without any error, confirming the correctness of the algorithm.

The Kosaraju's algorithm is a well-known algorithm for finding the strongly connected components (SCCs) of a directed graph. Here is a step-by-step description of the process:

1. **Reverse the Graph**: Create a new graph by reversing the direction of all edges in the original graph. This graph is called the transpose of the original graph.

2. **Run Depth-First Search (DFS) on the Original Graph**: Perform a DFS on the original graph, keeping track of the finish times of each vertex. The finish time is the time when a vertex has finished processing all its descendants.

3. **Order Vertices by Decreasing Finish Time**: After completing the DFS, sort all vertices in decreasing order of their finish times.

4. **Run DFS on the Transpose Graph**: Using the vertices in the order obtained in the previous step, perform a DFS on the transpose graph. Each DFS tree in this phase will represent a strongly connected component.

5. **Collect SCCs**: The vertices visited during each DFS in the transpose graph form a strongly connected component.

### Detailed Example Using the Image Provided:

The image shows an example with vertices and edges, labeled with numbers and finish times:

1. **Original Graph**:

   - Vertices: {1, 2, 3, 4, 5, 6, 7, 8, 9}
   - Edges: (between the vertices as shown)

2. **Step 1: Reverse the Graph**:

   - Reverse all edges of the original graph to create the transpose graph.

3. **Step 2: Run DFS on the Original Graph**:

   - Perform DFS and assign finish times to each vertex.
   - Example of finish times: \(f(1)=7, f(2)=4, f(3)=1, f(4)=8, f(5)=2, f(6)=5, f(7)=9, f(8)=6, f(9)=3\).

4. **Step 3: Order Vertices by Decreasing Finish Time**:

   - Order vertices based on finish times: 7, 4, 8, 6, 1, 9, 2, 5, 3.

5. **Step 4: Run DFS on the Transpose Graph**:

   - Using the ordered list, perform DFS on the transpose graph.
   - Each DFS tree represents a strongly connected component.
   - Example leaders (SCCs): Leader = 7, Leader = 6, Leader = 4.

6. **Step 5: Collect SCCs**:
   - Group vertices based on the DFS trees in the transpose graph.
   - Example SCCs: {7}, {6, 8, 9}, {4, 5, 3, 2}.

This process will identify all the strongly connected components in the graph.

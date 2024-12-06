Quicksort is a popular sorting algorithm that is based on the divide and conquer strategy. It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

Here’s an improved explanation of Quicksort:

## How Quicksort Works:

1. **Selecting a Pivot**: The choice of pivot is crucial for the efficiency of Quicksort. There are various strategies like choosing the first element, the last element, the median, or a random element as the pivot.

2. **Partitioning**: After choosing a pivot, the array is partitioned into two sub-arrays - elements less than the pivot and elements greater than the pivot. The pivot is then placed in its correct position in the sorted array.

3. **Recursion**: Quicksort is then recursively applied to the sub-arrays. This process continues until the base case is reached, which is an array with one or zero elements.

## Performance Analysis:

- **Worst-Case Scenario**: If the smallest or largest element is consistently chosen as the pivot, the algorithm will perform poorly with O(n^2) time complexity. This situation can occur if the array is already sorted or reverse sorted.

- **Best-Case Scenario**: The best case occurs when the pivot divides the array into two equal halves leading to a time complexity of O(n log n).

## Pivot Selection Strategies:

- **First or Last Element**: This strategy can lead to the worst-case time complexity, especially if the array is already sorted.

- **Median**: Ideally, choosing the median element as the pivot results in equal-sized partitions. If this strategy is consistently successful, the algorithm runs in O(n log n) time, similar to merge sort.

- **Random Selection**: By randomly selecting a pivot, the algorithm avoids the pitfall of a worst-case scenario tied to the input's initial order. On average, a randomly chosen pivot yields O(n log n) time complexity, making it a reliable choice for most applications.

## Random Pivots vs. Median Pivots:

- **Random Pivot**: Simplifies the algorithm and typically yields good performance regardless of the input's initial order. It averages out to O(n log n) time complexity.

- **Median Pivot**: Potentially yields a perfectly balanced split, but finding the true median is computationally intensive. However, using a 'median of three' (where the median is chosen among the first, middle, and last elements) strategy can be a good compromise.

## Conclusion:

Quicksort is generally faster than other O(n log n) algorithms like merge sort, due to lower overhead and better locality of reference. However, its performance heavily depends on the choice of pivot. In practical implementations, a hybrid approach often works best, such as using insertion sort for small sub-arrays and a 'median of three' for pivot selection to balance the efficiency and computational cost. Random pivots are a good default choice due to their simplicity and good average case performance.

## when chosen the first element of the array

![](https://storage.googleapis.com/zenn-user-upload/b7b628769612-20240429.png)

![](https://storage.googleapis.com/zenn-user-upload/9c24e90ad398-20240429.png)

## reference

https://www.youtube.com/watch?v=ETo1cpLN7kk&list=PLXFMmlk03Dt7Q0xr1PIAriY5623cKiH7V&index=26
https://github.com/subaru-hello/stanford-algorithms-go/blob/main/Course1/quick-sort/index.go
https://github.com/pco2699/algorithms/blob/master/quicksort/src/main.rs

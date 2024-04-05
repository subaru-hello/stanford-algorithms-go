## Merge Sort

Merge Sort is a quintessential example of the Divide and Conquer algorithmic paradigm. Initially, it divides a given integer array into two halves until the size of any subdivided array becomes one or empty. Subsequently, it sorts each pair of subarrays and merges them back together. A key operation during the merging phase is counting the inversions between two sorted arrays.

## Base Case

The base case occurs when the array's size is one or zero, at which point the array itself is returned. It's worth noting that the integers within the array are distinct.

## Efficiency of Merge Sort Algorithm

The Merge Sort algorithm ensures that every integer in the left-hand side array is smaller than those on the right-hand side, through iterative comparison.

Given:

- `n`: the size of the input array.
- `C`: the output array.

And considering:

- `A`: the first sorted half of the array.
- `B`: the second sorted half of the array.
- `i` and `j` starting at the first element.

The primary loop iterates from `k` to `n`, with conditions to merge `A` and `B` into `C`. Here, the algorithm's running time is influenced by the number of operations performed:

```
For each element in the for-loop, there are 4 operations, plus 2 initializations for `i` and `j`. This implies that the operation count directly correlates with the array size `m`, leading to an O(n) complexity.
```

However, since `A` and `B` are pre-sorted before merging, an additional 2 operations are considered before the loop, adjusting the count to 6m, excluding the initial 4m + 2. As `m` is halved in subsequent recursions, the total operations can be described by the formula 6m log2(m) + 6m, where +6m accounts for the operational baseline in each recursion. Log2(m) represents the depth of the recursion tree, indicating that the total operations at level `j` are <= 2^j \* 6(n/2^j) = 6n.

Thus, Merge Sort operates within a time complexity of 6m log2(m) + 6m, illustrating that, for an array of size one, a minimum of 6 operations are needed to complete the sorting. This model underscores the efficiency and scalability of Merge Sort in handling large datasets.

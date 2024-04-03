## merge-sort

the merge sort is a one of the Divide and conquar programming.
At first, it divides a given integer array into two parts unless the size of array reach under one or empty.
Next, it sorts each of pair subarrays and merge them back together.
Finally, count the inversions while merging two sorted arrays.

## base case

if array size lowers one or zero, just return array
integers in array are unique

## why does the the merge sort algorithms not surpass 6m log2^m + 6m

iterate two arrays to reach every left side array will be shorter than right side.
b = size
i = 1
j = 1

if(b[i] > b[j])
b[i] = i
i++
elsif()
b[j] = j
j++

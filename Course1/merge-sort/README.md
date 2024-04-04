## merge-sort

the merge sort is a one of the Divide and conquar programming.
At first, it divides a given integer array into two parts unless the size of array reach under one or empty.
Next, it sorts each of pair subarrays and merge them back together.
Finally, count the inversions while merging two sorted arrays.

## base case

if array size lowers one or zero, just return array
integers in array are unique

## why does the the merge sort algorithms not surpass 6m log2^m + 6m

iterate over two arrays to reach every left hand side integer of array will be shorter than right hand side.

```
n = inputed array size
C = output array

A = 1st sorted array
B = 2nd sorted array
i = 1
j = 1

first for loop using k to n
    if(A[i] < B[j])
        C[k] = A[i]
        i++
    elsif(B[j] < A[i])
        C[k] = B[j]
        j++
    end
```

basically, a runnning time would be calculated by a line of code executed,
so, if it required to estimate how many times to finish executing code, just count operations of the algorithms code step.

then, use the above code to get a running time formula.

the number of running times will be 4m + 2 as the following annotation "1" sum.
each "for loop" element has 4 operations and fixed operatioins are i and j.
as the size of array m increasing, for operations will eventually increasing O(n).

```
n = inputed array size
C = output array(n)

A = 1st sorted array(2/n)
B = 2nd sorted array(2/n)
i = 1st element // ←１
j = 1st element // ←1

first for loop using k to n
    if(A[i] < B[j]) // ←１
        C[k] = A[i] // ←１
        i++
    elsif(B[j] < A[i]) // ←１
        C[k] = B[j] // ←１
        j++
    end
```

but, A and B had already been sorted before merge function, it should be addded 2 operations before executing for loop.
So, it will be only 6m not 4m + 2.
And, m will be diveded by two after one executions. it means the size of m decreases unless it reaches one or zero.
Thus, it should be 6mlog2^m + 6m, +6m means how much operations does one execution have.
Merge Sort requires 6mlog2^m + 6m.
let's say the size is one, 0 + 6 = 6 , 6 operations are required to finish merge sort.
log2^m is the recursion tree depth.
total operations of the level j <= 2^j \* 6(n/2^j) = 6n

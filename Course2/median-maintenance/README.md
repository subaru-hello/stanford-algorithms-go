## Overview

Median maintenance is a problem that is used to find the median of a stream of numbers.

The Heap data structure could solve this problem given that Heap is a complete binary tree and it has the property that the parent node is always smaller than its children.

The time complexity of the algorithm is O(log n) because the size n will decrease log(n) in each iteration.

The space complexity is O(n) because we are using a Heap to store the numbers.

## Initial Setup

- given a list of numbers, x1, x2, x3 ...xn, one by one
- you need to tell the median of the numbers after each insertion

## Constraints

- The algorithm should work in O(log i) time at each step i.

## Solution

- use two heaps to maintain the median.
  - A max heap Hlow maintain the smaller half of the numbers.
  - A min heap Hhigh maintain the larger half of the numbers.

## Key Idea

- Maintain the invariant the size of Hlow is either equal to or 1 more than that of High.
- The max-heap Hlow supports extracting the maximum element.
- The min-heap Hhigh supports extracting the minimum element.

## Insertion Step

- When a new number x is added:
  1. Compare x with the maximum number of Hlow:
     - If x is smaller than or equal to the maximum of Hlow, insert x into Hlow.
     - Otherwise, insert x into Hhigh
  2. Balance the heaps
  - If the size of Hlow exceeds the size of Hhigh by more than 1, move the maximum element from Hlow to Hhigh.
  - If the size of Hhigh exceeds the size of Hlow by more than 1, move the minimum element from Hhigh to Hlow.

## Median Calculation

- If the sizes of the two heaps are equal, the median is the average of the maximum element of Hlow and the minimum element of Hhigh.
- If Hlow has one more element than Hhigh, the median is the maximum element of Hlow

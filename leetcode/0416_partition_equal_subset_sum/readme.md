# [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)

## Description
>Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
>
>Note:
> 1. Each of the array element will not exceed 100.
> 2. The array size will not exceed 200.

Example:

```
Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
```

## Solution

此问题可以使用0-1背包问题的思路求解

c是数组和的一半。

状态定义：F(n c)

状态转移方程：F(n c) = max(    F(n-1, c) ,     n + F(n-1, c-n)          )


[code](./partition_equal_subset_sum.go)

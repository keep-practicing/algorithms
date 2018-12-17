# 0-1背包问题

## 问题描述

```
有一个背包，它的容量为C(Capacity)，现在有n种不同的物品，编号为0...n-1，其中每一件物品的重量为w(i)，价值为v(i)。问可以向背包中盛放哪些物品，使得在不超过背包容量的基础上，物品的总价值最大。
```

## 解题思路

```
这类组合问题，可以使用递归完成，如果在其中能找到重叠子问题，就可以转化为记忆化搜索或动态规划来解决。

约束条件：
1. 要在n个物品里面选。
2. 容量小于等于一个给定的值C。
```

## 状态定义

```
F(n, C) 考虑将n个物品放进容量为C的背包，使得价值最大。
```

### 状态转移方程
```
F(i, c) = max(F(i-1, c), v(i)+F(i-1, c-w(i)))
```

### 递归实现
 ```go
func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}
	return bestValue(w, v, len(w)-1, c)
}

func bestValue(w []int, v []int, index int, c int) int {
	// 用[0...index]  的物品，填充容积为c的背包的最大价值。
	if index < 0 || c <= 0 {
		return 0
	}

	res := bestValue(w, v, index-1, c)
	if c >= w[index] {
		temp := v[index] + bestValue(w, v, index-1, c-w[index])
		if temp > res {
			res = temp
		}
	}
	return res
}
 ```

 ### 记忆化搜索
> 递归求解，index和c构成的数据对，存在重叠。

```go
func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}

	memo := make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		for j := 0; j < c+1; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	return bestValue(w, v, len(w)-1, c, memo)
}

func bestValue(w []int, v []int, index int, c int, memo [][]int) int {
	// 用[0...index]  的物品，填充容积为c的背包的最大价值。
	if index < 0 || c <= 0 {
		return 0
	}

	if memo[index][c] != -1 {
		return memo[index][c]
	}

	res := bestValue(w, v, index-1, c, memo)
	if c >= w[index] {
		temp := v[index] + bestValue(w, v, index-1, c-w[index], memo)
		if temp > res {
			res = temp
		}
	}
	memo[index][c] = res
	return res
}
```

### 自底向上动态规划

* 分析问题

    背包容量为5。

    物品描述

    |id|0|1|2|
    |:--:|:--:|:--:|:--:|
    |weight|1|2|3|
    |value|6|10|12|

    动态规划过程

    |物品编号/背包容量|0|1|2|3|4|5|
    |:--:|:--:|:--:|:--:|:--:|:--:|:--:|
    |0|0|6|6|6|6|6|
    |1|0|6|10|16|16|16|
    |2|0|6|10|16|18|22|

* 代码实现
  
```go
func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}

	n := len(w)

	if n == 0 {
		return 0
	}

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < c+1; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	// 只把第0个物品放入背包，动态规划二维数组值的变化。循环变量i表示背包的容量。
	for i := 0; i <= c; i++ {
		if i >= w[0] {
			memo[0][i] = v[0]
		} else {
			memo[0][i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= c; j++ {
			memo[i][j] = memo[i-1][j]
			if j >= w[i] {
				temp := v[i] + memo[i-1][j-w[i]]
				if temp > memo[i][j] {
					memo[i][j] = temp
				}
			}
		}
	}
	return memo[n-1][c]
}
```

### 0-1背包问题优化

分析
```
上述实现时间复杂度O(n*c), 空间复杂度O(n*c).

通过状态转移方程F(i, c) = max(F(i-1, c), v(i)+F(i-1, c-w(i)))，第i行元素只依赖第i-1行元素。理论上，只需要保持两行元素，即空间复杂度：O(2*c) = O(c).
```

```go 
func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}

	n := len(w)

	if n == 0 {
		return 0
	}

	memo := make([][]int, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j < c+1; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	// 只把第0个物品放入背包，动态规划二维数组值的变化。循环变量i表示背包的容量。
	for i := 0; i <= c; i++ {
		if i >= w[0] {
			memo[0][i] = v[0]
		} else {
			memo[0][i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= c; j++ {
			memo[i%2][j] = memo[(i-1)%2][j] // 将逻辑地址，转为物理空间地址。
			if j >= w[i] {
				temp := v[i] + memo[(i-1)%2][j-w[i]]
				if temp > memo[i%2][j] {
					memo[i%2][j] = temp
				}
			}
		}
	}
	return memo[(n-1)%2][c]
}
```

### 0-1背包问题进一步优化
```
继续优化空间，在只使用一行空间的情况下，即使用一维数组，完成动态规划，数组更新从右向左进行。
时间复杂度是常数级别的优化。
```

```go 
func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}

	n := len(w)

	if n == 0 {
		return 0
	}

	memo := make([]int, c+1)
	for i := 0; i < c+1; i++ {
		memo[i] = -1
	}

	// 只把第0个物品放入背包，动态规划二维数组值的变化。循环变量i表示背包的容量。
	for i := 0; i <= c; i++ {
		if i >= w[0] {
			memo[i] = v[0]
		} else {
			memo[i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := c; j >= w[i]; j-- {
			temp := v[i] + memo[j-w[i]]
			if temp > memo[j] {
				memo[j] = temp
			}
		}
	}
	return memo[c]
}
```

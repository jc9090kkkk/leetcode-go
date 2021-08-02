// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。

// 示例 1：
//
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶

// 解题思路：
// 通过题目可推断，1个台阶有1种方法，2个台阶有2种方法，3个台阶有3种方法，4个台阶有5种方法，5个台阶有5种方法，6个台阶有13种方法...
// 依次类推，可以发现题目的题解规律其实就是斐波那契数列
// 根据斐波那契数列的公式
// 时间复杂度为O(log n)
// 空间复杂度为O(1)

package main

import (
	"fmt"
	"math"
)

func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}

func main() {
	fmt.Println(climbStairs(44))
}
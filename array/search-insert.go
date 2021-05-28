// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 你可以假设数组中无重复元素。
//
// 示例 1:
// 输入: [1,3,5,6], 5
// 输出: 2

// 示例 2:
// 输入: [1,3,5,6], 2
// 输出: 1

// 示例 3:
// 输入: [1,3,5,6], 7
// 输出: 4

// 示例 4:
// 输入: [1,3,5,6], 0
// 输出: 0

// 解题思路：
// 1，数组是有序，一般对有序数组做操作基本都可以使用二分法，二分法的思想就是分而治之，通过不断计算中位数来减少判断的范围，最后再结合相应的判断条件，达到题目的要求
// 2，常规的二分法肯定无法满足题目要求，所以需要在二分的基础上调整下判断条件

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 设定左右边界 对应的其实就是数组下标
	left, right := 0, len(nums) - 1

	// 默认的索引位，这个位置不能为0，思考下为什么？
	//【因为如果target的值大于数组的最大值，index应该为数组最大下标+1，也就是len(nums)，如果是0的话，返回值就是错的】
	index := len(nums)
	// 左边界 不大于 有边界 二分法判断才有意义
	for left <= right {
		// 计算中位
		middle := (right + left) / 2

		// 根据判断条件 重新定义左右边界 并更新当前索引值
		if target <= nums[middle] {
			index = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	// 返回索引值
	return index
}

func main() {
	arr := []int{1, 3, 5, 6}
	res := searchInsert(arr, 2)
	fmt.Printf("元素索引值为:%d", res)
}
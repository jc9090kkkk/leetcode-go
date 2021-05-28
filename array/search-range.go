// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]，请用时间复杂度为O(log n)的算法实现。
//
// 示例 1：
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

// 示例 2：
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]

// 示例 3：
// 输入：nums = [], target = 0
// 输出：[-1,-1]

// 解题思路：
// 数组为有序数组，题目要求时间复杂度为O(log n)，所以可以使用二分查找
// 如果先找到中间值，然后再依次向左右两侧展开查找，如果只用一个二分查询，会陷入死循环，因为左右两侧的边界随着middle计算，无法跳出最后的循环条件
// 所以要换另外一种方式，用两个二分查找分别查询同一个数组，其中一个计算元素的最左侧索引位置，也就是第1个位置
// 另外一个计算元素的最右侧索引位置，也就是最后一个位置，最后再合起来，两个同等的二分法时间复杂度依旧是O(log n)，满足题目要求

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	// 先设定元素开始位置和结束位置
	position := []int{-1, -1}

	// 查询出目标值的左侧索引
	leftIndex := binaryLeftSearch(nums, target)
	// 查询出目标值的右侧索引
	rightIndex := binaryRightSearch(nums, target) - 1

	// 返回最终索引切片数据
	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		position = []int{leftIndex, rightIndex}
	}

	return position
}

// 二分左侧查询
func binaryLeftSearch(nums []int, target int) int {
	ans := len(nums)

	// 设定左右边界
	left := 0
	right := ans - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] >= target {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

// 二分右侧查询
func binaryRightSearch(nums []int, target int) int {
	ans := len(nums)

	// 设定左右边界
	left := 0
	right := ans - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

func main() {
	arr := []int{5, 7, 7, 8, 8, 8, 8, 10}
	res := searchRange(arr, 8)
	fmt.Println(res)
}

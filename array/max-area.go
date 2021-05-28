// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
// 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。

// 示例:
// 输入:[1, 8, 6, 2, 5, 4, 8, 3, 7]
// 输出: 49

// 解题思路：
// 这道题最大面积如何保证，第一：横坐标系尽量要宽，第二：两个纵坐标中高度最短（元素值大小）的那一个决定了最大面积的最终结果，所以纵坐标也要尽量高，
// 其中横坐标的大小其实就是数组中左右两侧元素的索引值之差，纵坐标的高度就是数组元素自己本身，
// 所以最大面积的计算公式就是（rightIndex - leftIndex）* max(arr[right],arr[left])，利用双指针法，只需要一次遍历就可以计算出最大的盛水面积了，
// 这样时间复杂度为O(N), 空间复杂度为O(1)

package main

import "fmt"

func maxArea(height []int) int {
	// 定义左右指针
	left := 0
	right := len(height) - 1

	// 暂时定义最大面积为0
	maxArea := 0
	// 临时面积为0
	currentArea := 0
	// 左右指针中较小的值
	minPointer := 0

	for left < right {
		// 判断左右指针数值中相对较小的一个 用于面积计算
		if height[left] <= height[right] {
			minPointer = height[left]
		} else {
			minPointer = height[right]
		}

		// 计算当前条件下的最大面积 如果当前面积大于之前计算的最大面积 则替换
		currentArea = (right - left) * minPointer
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// 如果左侧元素小于右侧元素 则左指针向右移位 否则是右指针向左移位
		if height[left] < height[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}

	return maxArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(arr)
	fmt.Printf("最大盛水面积为:%d", res)
}

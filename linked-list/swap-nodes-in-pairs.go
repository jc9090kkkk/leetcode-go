// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
//
// 输入：head = []
// 输出：[]
//
// 输入：head = []
// 输出：[]

// 解题思路: 先定义一个零时的头部节点，用这个头部节点去和现在已有的链表节点的Next指针做交换，交换完后启示位置跳到已经交换过的节点作为头节点
// 继续下一轮交换 直到所有的节点都遍历一遍
// 时间复杂度:O(N)
// 空间复杂度:O(1)
// 题解链接:https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-91/

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapNodesInPairs(head *ListNode) *ListNode {
	// 先定义一个头部节点
	headNode := &ListNode{0, head}
	temp := headNode
	// 持续判断后续的第一个node和第二个node是否为空 不为空则进行交换
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	return headNode.Next
}

func main() {
	linkedList4 := ListNode{4, nil}
	linkedList3 := ListNode{3, &linkedList4}
	linkedList2 := ListNode{2, &linkedList3}
	linkedList := ListNode{1, &linkedList2}

	swapLinkedList := swapNodesInPairs(&linkedList)
	swapLinkedList2 := swapNodesInPairs(&linkedList2)
	swapLinkedList3 := swapNodesInPairs(&linkedList3)
	swapLinkedList4 := swapNodesInPairs(&linkedList4)

	fmt.Println(swapLinkedList.Val)
	fmt.Println(swapLinkedList2.Val)
	fmt.Println(swapLinkedList3.Val)
	fmt.Println(swapLinkedList4.Val)

}


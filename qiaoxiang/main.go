package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{2, 1, 4, 3, 7, 5, 3}
	///QuickSort(nums)
	MergeSort(nums)
	fmt.Println(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	cur := head

	for i := 1; i < k; i++ {
		cur = cur.Next
		if cur == nil {
			return head
		}
	}
	nextHead := cur.Next
	tail := reverseList(head, cur)
	tail.Next = reverseKGroup(nextHead, k)
	return cur
}

func reverseList(head, tail *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre != tail {
		cur = head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	for pre.Next != nil {
		pre = pre.Next
	}
	return pre
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	var mostRight *TreeNode
	var res []int
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

//62. 不同路径
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	f := make([][]int, m, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = 1
				continue
			}
			if i > 0 {
				f[i][j] += f[i-1][j]
			}
			if j > 0 {
				f[i][j] += f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}

//63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	f := make([][]int, m, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				f[i][j] = 1
			}

			if i > 0 {
				f[i][j] += f[i-1][j]
			}
			if j > 0 {
				f[i][j] += f[i][j-1]
			}

		}
	}
	return f[m-1][n-1]
}

// QuickSort 快速排序
func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
}

func partition(nums []int, l, r int) (int, int) {
	lt := l
	gt := r + 1
	d := rand.Intn(r-l+1) + l
	nums[l], nums[d] = nums[d], nums[l]
	v := nums[l]
	k := l + 1
	for k < gt {
		if nums[k] < v {
			nums[lt+1], nums[k] = nums[k], nums[lt+1]
			lt++
			k++
		} else if nums[k] > v {
			nums[gt-1], nums[k] = nums[k], nums[gt-1]
			gt--
		} else {
			k++
		}
	}
	nums[lt], nums[l] = nums[l], nums[lt]
	return lt - 1, gt
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	lt, gt := partition(nums, l, r)
	quickSort(nums, l, lt)
	quickSort(nums, gt, r)
}

// MergeSort 归并排序
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	buf := make([]int, r-l+1, r-l+1)
	for i := l; i <= r; i++ {
		buf[i-l] = nums[i]
	}
	i := 0
	mid = mid - l
	j := mid + 1
	for l <= r {
		if i > mid {
			nums[l] = buf[j]
			j++
		} else if j >= len(buf) {
			nums[l] = buf[i]
			i++
		} else if buf[j] < buf[i] {
			nums[l] = buf[j]
			j++
		} else {
			nums[l] = buf[i]
			i++
		}
		l++
	}
}

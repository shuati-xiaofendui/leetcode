package main

func main() {

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

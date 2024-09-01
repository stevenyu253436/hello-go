package main

import "fmt"

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 将排序数组转换为高度平衡的二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	// 如果数组为空，返回 nil
	if len(nums) == 0 {
		return nil
	}
	return helper(nums, 0, len(nums)-1)
}

// 辅助函数，递归构建树
func helper(nums []int, left int, right int) *TreeNode {
	// 递归终止条件
	if left > right {
		return nil
	}

	// 选择中间元素作为根节点
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	// 递归构建左子树和右子树
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

// 测试代码
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Println(tree)
}

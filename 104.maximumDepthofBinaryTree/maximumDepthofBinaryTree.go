package maximumDepthofBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *	  		 3
 * 			/ \
 *		   9  20
 *			 /  \
 *		    15   7
 *
 *  [3,9,20,null,null,15,7],
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 一层一层返回值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)  // => 节点9 => root=nil 返回 0+1 = 1
	r := maxDepth(root.Right) // => 节点20 => 左节点15(=> root=nil => 返回0+1=1)；右节点7(=> root=nil => 返回0+1=1) => 返回1+1=2
	return max(r, l) + 1      // 最后是 l = 1 r = 2 => 返回 3

}

func max(r, l int) int {
	if r > l {
		return r
	}
	return l
}

package validatebst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBST(root *TreeNode, value *int) bool {
	if root == nil {
		return true
	}
	if !isBST(root.Left, value) {
		return false
	}
	if *value >= root.Val {
		return false
	}
	*value = root.Val
	return isBST(root.Right, value)
}

func isValidBST(root *TreeNode) bool {
	value := math.MinInt
	return isBST(root, &value)
}

//  func isValidBST(root *TreeNode) bool {
// 	return validate(root, math.MinInt, math.MaxInt)
// }

// func validate(node *TreeNode, low int, high int) bool {
// 	if node == nil {
// 		return true
// 	}

// 	if  low < node.Val && node.Val < high {
// 	    return validate(node.Left, low, node.Val) && validate(node.Right, node.Val, high)
// 	}
//     return false

// }

// func min(i int, j int) int {
//     if i>=j {
//         return j
//     }
//     return i
// }

// func max(i int, j int) int {
//     if i>=j {
//         return i
//     }
//     return j
// }
// func minValueNode(root *TreeNode) int {
//     if (root == nil) {
//         return math.MaxInt32
//     }
//     return min(root.Val, min(minValueNode(root.Left),minValueNode(root.Right)))
// }

// func maxValueNode(root *TreeNode) int {
//     if (root == nil) {
//         return math.MinInt32
//     }
//     return max(root.Val, max(maxValueNode(root.Left),maxValueNode(root.Right)))
// }

// func isValidBST(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }
//     if root.Left != nil && root.Val <= maxValueNode(root.Left) {
//         return false
//     }
//     if root.Right != nil && root.Val >= minValueNode(root.Right) {
//         return false
//     }
//     return isValidBST(root.Left) && isValidBST(root.Right)
// }

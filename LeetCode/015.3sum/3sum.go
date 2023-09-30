package sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; {
		x := i + 1
		y := len(nums) - 1
		for x < y {
			sum := nums[i] + nums[x] + nums[y]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[x], nums[y]})
			}
			if sum > 0 {
				for y = y - 1; x < y && nums[y] == nums[y+1]; y-- {
				}
			} else {
				for x = x + 1; x < y && nums[x] == nums[x-1]; x++ {
				}
			}
		}
		for i = i + 1; i < len(nums) && nums[i-1] == nums[i]; i++ {
		}
	}

	return res
}

// func threeSum(nums []int) [][]int {
// 	count := 1
// 	res := [][]int{}
// 	sort.Ints(nums)
// 	for i := 1; i < len(nums); {
// 		if nums[i] == nums[i-1] {
// 			if count == 2 {
// 				nums = append(nums[:i], nums[i+1:]...)
// 				if nums[i-1] == 0 {
// 					if len(res) == 0 {
// 						res = append(res, []int{0, 0, 0})
// 					}
// 					nums = append(nums[:i-1], nums[i:]...)
// 					i--
// 					count--
// 				}
// 			} else {
// 				count++
// 				i++
// 			}
// 		} else {
// 			count = 1
// 			i++
// 		}

// 	}
// 	try(nums, &res)
// 	return res
// }

// func try(nums []int, res *[][]int) {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			rest := -nums[i] - nums[j]
// 			index := sort.SearchInts(nums[j+1:], rest)

//				if index < len(nums[j+1:]) && nums[index+j+1] == rest {
//					check := true
//					for x := 0; x < len(*res); x++ {
//						if (*res)[x][0] == nums[i] && (*res)[x][1] == nums[j] && (*res)[x][2] == rest {
//							check = false
//						}
//					}
//					if check {
//						*res = append(*res, []int{nums[i], nums[j], rest})
//					}
//				}
//			}
//		}
//	}
func Main() {
	// var nums = []int{-1, 0, 1, 2, -1, -4}
	var nums = []int{-2, 0, 1, 1, 2}
	// var nums = []int{1, -1, -1, 0}
	// nums := []int{3, 0, -2, -1, 1, 2}
	// nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	fmt.Println(len(nums))
	fmt.Println(threeSum(nums))

}

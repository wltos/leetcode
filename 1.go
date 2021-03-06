/*
[题目]
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

[举例]
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

[思路]
目标值 - 数组元素，然后挨个比对，查找差值是否在数组中，最后返回减数和差的下标
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 4, 6}
	fmt.Println(a)
	fmt.Println(twoSum(a, 10)) // 目标值是10
}

func twoSum(nums []int, target int) []int {
	for i := range nums {
		data := target - nums[i]
		for j := range nums {
			if data == nums[j] && j != i {
				return []int{i, j}
			}
		}
	}
	return nil
}

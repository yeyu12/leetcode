package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// 思路，把数组转为map结构，在转的过程中，查找target-v是否在numsMap结构中存在，如果存在则表示已经找到，如果未找到则把kv存储到map中。如果一直到循环完都未找到，则直接返回空数组
// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		if i, ok := numsMap[target-v]; ok {
			return []int{i, k}
		} else {
			numsMap[v] = k
		}
	}

	return []int{}
}

// @lc code=end

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

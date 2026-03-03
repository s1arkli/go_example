package main

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
示例 1 ：

输入：nums = [2,2,1]
*/

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ch := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ch[nums[i]]++
	}
	for k, v := range ch {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 按位异或
func singleNumber_(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}

//相同的异或=0，初始为0，最终=0^single=single

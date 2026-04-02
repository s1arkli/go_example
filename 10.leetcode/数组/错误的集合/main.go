package main

/*
集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。

给定一个数组 nums 代表了该集合发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
*/

// todo 对于题目所提供的数组，每个数都是有对应的位置的，对对应下标取反做标记，重复标记的数即为重复数。未被标记的下标即是缺失的数所映射的下标，通过这个下标即可计算出缺失的数
func findErrorNums(nums []int) []int {
	dup, miss := 0, 0
	//使用下标对应，由于只重复了一个数，所以通过值映射下标一定有重复映射的
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			dup = abs(nums[i])
		}
	}
	//一定有未映射到的，未映射到的下标即为缺失的数字
	for k, num := range nums {
		if num > 0 {
			miss = k + 1
		}
	}
	return []int{dup, miss}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//对于全是正数切不重复的数组，对值取反也是一种标记手段

package main

/*
给定一个整数数组  nums，处理以下类型的多个查询:

计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int left, int right) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
*/

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{
		arr: prefix,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.arr[right+1] - this.arr[left]
}

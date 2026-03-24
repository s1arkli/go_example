package main

func main() {

}

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}

		digits[i] = 0

		if i == 0 {
			return append([]int{1}, digits...)
		}
	}
	return digits
}

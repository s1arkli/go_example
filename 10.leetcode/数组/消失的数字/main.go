package main

func missingNumber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	seen := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = true
	}
	for i := 0; i <= len(nums); i++ {
		_, ok := seen[i]
		if !ok {
			return i
		}
	}

	return 0
}

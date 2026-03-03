package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	lastIdx := make(map[byte]int)
	maxLen, left := 0, 0

	for right := 0; right < len(s); right++ {
		bt := s[right]
		if idx, exist := lastIdx[bt]; exist && idx >= left {
			left = idx + 1
		}
		lastIdx[bt] = right
		if length := right - left + 1; length > maxLen {
			maxLen = length
		}

	}

	return maxLen
}

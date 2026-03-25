package main

/*
premium lock icon
相关企业
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
*/

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	need, window := [26]int{}, [26]int{}
	left := 0
	res := make([]int, 0)

	for _, v := range p {
		//v-a把下标压缩到26以内，++代表出现次数+1
		need[v-'a']++
	}

	for right := 0; right < len(s); right++ {
		window[s[right]-'a']++

		if right-left+1 > len(p) {
			//超出了滑动窗口，left++
			window[s[left]-'a']--
			left++
		}

		if need == window {
			res = append(res, left)
		}
	}
	return res
}

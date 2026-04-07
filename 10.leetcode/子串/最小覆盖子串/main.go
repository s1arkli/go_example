package main

/*
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。
如果没有这样的子串，返回空字符串 ""。
*/

func minWindow(s string, t string) string {
	//使用滑动窗口完成
	nt := len(t)
	ns := len(s)
	valid := 0
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < nt; i++ {
		need[t[i]]++
	}
	validLen := len(need)

	left := 0
	start, end := 0, -1

	for right := 0; right < ns; right++ {
		c := s[right]
		if _, ok := need[c]; ok {
			//存到window内，并判断是否窗口内的该字节符合要求
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}

		for valid == validLen {
			if end == -1 {
				start = left
				end = right
			} else {
				if right-left < end-start {
					start = left
					end = right
				}
			}
			lc := s[left]
			if _, ok := need[lc]; ok {
				if window[lc] == need[lc] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}

	if end == -1 {
		return ""
	}

	return s[start : end+1]
}

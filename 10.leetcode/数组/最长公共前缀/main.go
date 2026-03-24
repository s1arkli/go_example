package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
*/

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

/*
1.我最初的想法是挨个对比每个字符串的前几位，但是我没想到怎么同时拿到每个元素的第一位

2.或者拿到第一个的字母和后面做比较
*/

//func longestCommonPrefix(strs []string) string {
//	res := ""
//	yes := true
//	for i := 0; i < len(strs[0]); i++ {
//		pre := strs[0][i]
//
//		for j := 1; j < len(strs); j++ {
//			if pre == strs[j][i] {
//				continue
//			}
//			if pre != strs[j][i] {
//				yes = false
//			}
//		}
//		if yes {
//			res = string(append([]byte(res), pre))
//		} else {
//			break
//		}
//	}
//	return res
//}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		//对首个元素遍历，然后和后续元素进行对比
		for j := 1; j < len(strs); j++ {

			if i >= len(strs[j]) {
				//这是左闭右开，搞错成[:i-1]
				return strs[0][:i]
			}

			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

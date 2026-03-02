package main

import (
	"sort"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

func groupAnagrams(strs []string) [][]string {
	strMp := make(map[string][]string, len(strs))
	//将输入按照字母进行分类，相同字母放到同一个key下
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		strMp[string(s)] = append(strMp[string(s)], str)
	}
	resp := make([][]string, 0, len(strMp))
	for _, strs := range strMp {
		resp = append(resp, strs)
	}
	return resp
}

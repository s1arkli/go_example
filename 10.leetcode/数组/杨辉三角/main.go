package main

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
*/

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	if numRows == 0 {
		return res
	}

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1

		if i > 1 {
			for j := 1; j < i; j++ {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}

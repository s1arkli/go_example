package main

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

tips:从第0行开始
*/

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	//在一个数组上面计算每一行杨辉三角，并且倒过来计算（这是由于正着算会把前面的元素提前改变导致后续计算无意义）
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			row[j] = row[j] + row[j-1]
		}
	}
	return row
}

package main

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

func main() {

}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	copy(nums1[m:], nums2)
//	sort.Ints(nums1)
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right := 0, 0
	res := make([]int, 0, m+n)

	for {
		if left == m {
			res = append(res, nums2[right:]...)
			break
		}

		if right == n {
			res = append(res, nums1[left:]...)
			break
		}

		if nums1[left] <= nums2[right] {
			res = append(res, nums1[left])
			left++
		} else {
			res = append(res, nums2[right])
			right++
		}
	}
	copy(nums1, res)
}

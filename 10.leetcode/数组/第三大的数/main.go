package main

import "fmt"

func main() {
	var a *int
	fmt.Println(a)
}

func thirdMax(nums []int) int {
	var (
		f, s, t *int
	)

	for _, v := range nums {
		n := v
		if f == nil || n > *f {
			t = s
			s = f
			f = &n
		} else if n < *f && (s == nil || n > *s) {
			t = s
			s = &n
		} else if s != nil && n < *s && (t == nil || n > *t) {
			t = &n
		}
	}

	if t == nil {
		return *f
	}
	return *t
}

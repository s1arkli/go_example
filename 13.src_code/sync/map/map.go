package _map

import (
	"fmt"
	"sync"
)

var tt sync.Map

func main() {
	defer tt.Clear()
	tt.Store("key1", "value1")
	tt.Store("key2", "value2")

	fmt.Println(tt.Load("key1"))
	fmt.Println(tt.Load("key2"))

	tt.Range(func(key, value interface{}) bool {
		fmt.Println(value)
		return true
	})

	tt.Swap("key1", "value1")
}

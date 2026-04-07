package main

import "sync"

func main() {
	var a sync.Once

	a.Do(func() {})
}

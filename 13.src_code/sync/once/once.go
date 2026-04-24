package once

import (
	"fmt"
	"sync"
)

var a sync.Once

func main() {
	//只会打印一次
	go once()
	go once()
	go once()
}

func once() {
	a.Do(func() {
		fmt.Println("-------")
	})
}

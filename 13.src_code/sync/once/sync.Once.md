
内部结构
```go
type Once struct {  
    _ noCopy  
	done atomic.Bool  
    m    Mutex  
}
```

使用方法
```go
var a sync.Once
a.Do(func(){})
```

内部维护了一个原子bool变量（cpu指令），每次调用Do(fn)的时候，首先加锁，内部读取bool状态，判断是否为false，如为true，执行fn，修改bool为false。在多个go routine同时执行的场景下，保证只会执行一次。



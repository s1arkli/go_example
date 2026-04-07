
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

内部维护了一个原子bool变量（cpu指令），当第一次执行fn的时候，首先加锁，内部读取bool状态，执行fn。完毕后修改bool状态，在多个go routine同时执行的场景下，保证只会执行一次。



# sql


```go
    func Open(driverName, dataSourceName string) (*DB, error)
```
通过drivername获取对应驱动，将驱动+dsn封装为db对象返回。
我理解的驱动：内部包括tcp协议（握手、挥手等），数据格式的转换（将mysql返回的结果转换成变量给用户）等等。

---

```go
func (db *DB) Exec(query string, args ...any) (Result, error)
```
内部使用tcp连接发送sql命令，并解析数据库返回的结果。

```go
conn := db.freeConn[last]
```
首先从闲置连接池内取连接，如果没有再重新建立连接。
```go
ci, err := db.connector.Connect(ctx)
```
使用驱动器内部封装的connect方法建立连接。

```go
func (db *DB) execDC(ctx context.Context, dc *driverConn, release func(error), query string, args []any) (res Result, err error)
```

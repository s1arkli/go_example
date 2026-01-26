# gorm
gorm本质上是对go语言标准库sql包的进一步封装，在其之上提供了方便的事务管理、链式查询api、模型映射，从而降低开发的复杂度。

# sql
标准库sql，是建立连接池，通过下载对应关系型数据库的驱动，使用驱动产生具体连接。1.定义了一些规范接口，例如db.Exec,db.Query用于与数据库进行交互。
2.建立连接池，让连接可复用，管理连接数量和存活时间等。3.提供基础的事务管理，db.BeginTx，tx.Rollback,tx.Commit。4.可以更换驱动就可以切换不同
的关系性数据库，迁移成本低。

## config
mysql包有三个config结构体dsn,mysql,gorm。其中dsn.config保存的是客户端的用户名、密码、地址等客户端的连接信息。mysql.config主要是保存mysql
服务的基础信息，包括驱动名、mysql版本信息、dsn（data source name）、连接池。gorm.Config主要是给连接mysql的前置准备的相关参数，也就是orm的
配置信息，比如SkipDefaultTransaction--在执行insert单条sql的时候也会开启一个事务，这个字段开启代表这个事务。

## mysql.Open
```go
func Open(dsn string) gorm.Dialector {
	dsnConf, _ := mysql.ParseDSN(dsn)
	return &Dialector{Config: &Config{DSN: dsn, DSNConfig: dsnConf}}
}
```
解析dsn，拿到addr，pwd，user等连接信息，并初始化config

## gorm.Open
解析opt，并构建连接前的初始化，最后开始连接mysql
```go
err = config.Dialector.Initialize(db)
```
此处的config.Dialector是mysql.Dialector，也就是在mysql解析dsn的时候就已经定义好如何连接数据库了。


## db.Exec
```go
    tx = db.getInstance()
```
这是每次执行Exec时的第一步，内部根据db.Clone字段来区分是否需要状态隔离。在grom.Open的时候，db.Clone被赋值为1，此后调用db.getInstance()
都会实例化一个新的*gorm.DB指针，函数内部复用了db的连接池和上下文等内容，重新创建statement结构体赋值给tx实例并返回，其中clone=0。这一步的作用
是意图分化的节点进行状态隔离（例如db.Where_1,db.Where_2，这二者的db就是状态隔离的 *gorm.DB，在链式调用的情况下（db.Where().Limit()...）
，由于clone=0，每次都是复用同一个 *gorm.DB，将query进行叠加，最终执行sql）

```go
    if strings.Contains(sql, "@") {
		clause.NamedExpr{SQL: sql, Vars: values}.Build(tx.Statement)
	} else {
		clause.Expr{SQL: sql, Vars: values}.Build(tx.Statement)
	}
```
Exec支持两种入参方式，@是命名参数，不依赖参数位置。而后者依赖位置进行入参如main.go。其中Build的作用是把sql+参数给整理成mysql驱动所规定的数据格式，
并写入缓冲区待用。

```go
 tx.callbacks.Raw().Execute(tx)
```
使用注册的回调函数执行

```go
func RegisterDefaultCallbacks(){
	...
	
    rawCallback := db.Callback().Raw()
    rawCallback.Register("gorm:raw", RawExec)
    rawCallback.Clauses = config.QueryClauses
}
```
所有的回调函数在callbacks包内的callbacks.go内注册。其中这三行是注册db.Exec执行的具体函数的位置。
```go
func RawExec(db *gorm.DB) {
	if db.Error == nil && !db.DryRun {
		result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
		if err != nil {
			db.AddError(err)
			return
		}

		db.RowsAffected, _ = result.RowsAffected()

		if db.Statement.Result != nil {
			db.Statement.Result.Result = result
			db.Statement.Result.RowsAffected = db.RowsAffected
		}
	}
}
```
这是内部真正执行的函数入口，内部从连接池中获取db连接，执行具体query，连接池的管理等等。

## db.Raw
```go
func (db *DB) Raw(sql string, values ...interface{}) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.SQL = strings.Builder{}

	if strings.Contains(sql, "@") {
		clause.NamedExpr{SQL: sql, Vars: values}.Build(tx.Statement)
	} else {
		clause.Expr{SQL: sql, Vars: values}.Build(tx.Statement)
	}
	return
}
```
db.Raw对比db.Exec只是少了回调函数的使用，Raw只是构建查询，不涉及具体执行。

## Build
```go
func (expr NamedExpr) Build(builder Builder) 
```
Build函数的作用是将sql转换为[]byte，并逐一写入缓冲区，处理sql中的？按照规则绑定参数。

## 总结
gorm包的作用使用数据库驱动建立数据库连接，并将函数语言的crud转换成对应的sql方言经tcp连接发送给数据库，并拿到返回值。
准备---->构建---->执行---->收尾

## tips
- 在返回值处进行初始化，就可以不return具体的变量，在函数内部赋值即可。

- 在执行sql的时候有一个统一处理err的方式，此方式不论最后成功与否，只根据err的状态判断连接的状态并做出相应的操作（例如：好连接放回连接池，坏连接关闭等等）
是一个对错误处理的好方式。
```go
defer func() {
		release(err)
	}()
```

- gorm底层执行代码还是调用的go官方sql包，在使用方面官方go包和gorm差别不大
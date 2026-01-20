# mysql
mysql包的作用是根据输入的mysql配置，连接mysql服务，将sql语言发送给mysql执行并将得到的数据返回给客户端，完数据库的crud。

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

## 总结
gorm包的作用是建立mysql连接，并将sql语句发送给mysql执行，最终拿到返回结果。

## tips
- 在返回值处进行初始化，就可以不return具体的变量，在函数内部赋值即可。
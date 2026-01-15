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
笑话opt，并构建连接前的初始化，最后开始连接mysql
```go
err = config.Dialector.Initialize(db)
```
此处的config.Dialector是mysql.Dialector，也就是在mysql解析dsn的时候就已经定义好如何连接数据库了
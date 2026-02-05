# 告诉make不是文件名，直接执行命令即可
.PHONY: proto server

# :后面的是依赖，会优先执行依赖的shell在执行本体
server: proto
	go run ./server/main.go

# go_out/go--grpc_out = .pb是设置生成代码的输出路径
proto:
	protoc --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
	--proto_path=./proto \
	proto/test.proto
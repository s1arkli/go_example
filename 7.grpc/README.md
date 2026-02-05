## prepare
- **brew install protobuf**//protocol buffers编译器
- **go install google.golang.org/protobuf/cmd/protoc-gen-go@latest**
- **go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest**

---

## bash
```bash
protoc --go_out=./pb --go_opt=paths=source_relative \
       --go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
       --proto_path=./proto \
       proto/test.proto
```

**逐个参数解释**

| 参数                                    | 作用                                                     |
|---------------------------------------|--------------------------------------------------------|
| `--go_out=./pb`                       | 消息类型代码（`test.pb.go`）输出到 `pb/` 目录                       |
| `--go_opt=paths=source_relative`      | 生成路径相对于 proto 文件位置，不按 go_package 创建目录                  |
| `--go-grpc_out=./pb`                  | gRPC 服务代码（`test_grpc.pb.go`）输出到 `pb/` 目录               |
| `--go-grpc_opt=paths=source_relative` | 同上                                                     |
| `--proto_path=./proto`                | 告诉 protoc 去哪里找 proto 文件（可简写为 `-I./proto`）              |
| `proto/test.proto`                    | 要编译的 proto 文件                                          |


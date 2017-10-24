# struct_first_upper
go:generate tools

进行自动处理的文件名格式为 *_struct.go

在需要处理的包中加入
```go
//go:generate struct_first_upper

```

再直接执行命令
```go
go generate
```

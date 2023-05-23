# B站 gRPC API Golang 封装

接口大部分来自于 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)

## grpc主机

B站客户端的grpc接口主机为以下服务器

> grpc.biliapi.net
>
> app.bilibili.com

### 安装

```bash
go get -u github.com/XiaoMiku01/bilibili-grpc-api-go
```

### 示例  

- [example.go](example/example.go)：示例代码

### 注意事项

- 和B站建立的gRPC连接可以复用
- 每次请求都需要重新生成`metadata`头部
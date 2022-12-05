# B站 gRPC API Golang 封装



## grpc主机

B站客户端的grpc接口主机为以下服务器

> grpc.biliapi.net
>
> app.bilibili.com

## grpc鉴权

需要在请求http头部中添加`access_key`，如下

```
authorization:identify_v1 {access_key}
```

## grpc头部

- [bilibili.metadata](bilibili/metadata)：客户端环境参数
- [bilibili.rpc](bilibili/rpc/status.proto)：响应错误信息

## 接口请求定义

*稍后补充*
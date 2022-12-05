package test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	onlineapi "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/playeronline/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"log"
	"testing"
	"time"
)

// init a grpc client
// 初始化一个grpc客户端 可以复用
var grpcClient *grpc.ClientConn

// connect to grpc server
// 连接grpc服务器
func init() {
	addr := "grpc.biliapi.net:443"
	creds := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS12,
	}))
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}
	var err error
	grpcClient, err = grpc.Dial(addr, creds, grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("BiliGRPC did not connect: %v", err)

	}
	log.Println("BiliGRPC connect success")
}

// TestGetOnline test get online player
// 使用grpc获取视频在线人数测试用例
func TestGrpcApi(t *testing.T) {
	onlineClient := onlineapi.NewPlayerOnlineClient(grpcClient)
	onlineReq := &onlineapi.PlayerOnlineReq{
		Aid:      390373397,
		Cid:      898386251,
		PlayOpen: true,
	}
	onlineResp, err := onlineClient.PlayerOnline(context.Background(), onlineReq)
	if err != nil {
		t.Error(err)
	}
	jsonString, err := json.Marshal(onlineResp)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonString))
}

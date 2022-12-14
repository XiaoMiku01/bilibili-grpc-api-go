package test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/metadata/locale"
	"github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/rpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	dynamicapi "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/dynamic/v2"
	onlineapi "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/playeronline/v1"
	bilimetadata "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/metadata"

	"github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/metadata/device"
	"github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/metadata/network"
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

// getBiliBiliMetaData get bilibili metadata
// 获取bilibili元数据 用于鉴权
func getBiliBiliMetaData(accessKey string) metadata.MD {
	device := &device.Device{
		MobiApp:  "android",
		Device:   "phone",
		Build:    6830300,
		Channel:  "bili",
		Buvid:    "XX82B818F96FB2F312B3A1BA44DB41892FF99",
		Platform: "android",
	}
	devicebin, _ := proto.Marshal(device)
	locale := &locale.Locale{
		Timezone: "Asia/Shanghai",
	}
	localebin, _ := proto.Marshal(locale)
	bilimetadata := &bilimetadata.Metadata{
		AccessKey: accessKey,
		MobiApp:   "android",
		Device:    "phone",
		Build:     6830300,
		Channel:   "bili",
		Buvid:     "XX82B818F96FB2F312B3A1BA44DB41892FF99",
		Platform:  "android",
	}
	bilimetadatabin, _ := proto.Marshal(bilimetadata)
	network := &network.Network{
		Type: network.NetworkType_WIFI,
	}
	networkbin, _ := proto.Marshal(network)
	md := metadata.Pairs(
		"x-bili-device-bin", string(devicebin),
		"x-bili-local-bin", string(localebin),
		"x-bili-metadata-bin", string(bilimetadatabin),
		"x-bili-network-bin", string(networkbin),
		"authorization", "identify_v1 "+accessKey,
	)
	return md
}

// TestGetDynamic test get dynamic with grpc api by access key
// 使用grpc获取关注动态测试用例 access key 鉴权
func TestGrpcApiWithAuthorize(t *testing.T) {
	accessKey := "your access key"
	md := getBiliBiliMetaData(accessKey)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	dynamicClient := dynamicapi.NewDynamicClient(grpcClient)
	dynAllReq := &dynamicapi.DynAllReq{
		Offset: "",
		Page:   1,
	}
	dynAllResp, err := dynamicClient.DynAll(ctx, dynAllReq)
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			t.Error(err)
			return
		}
		// B站的grpc接口返回的错误码 例如鉴权错误
		if status.Code() == codes.Unknown && len(status.Details()) > 0 {
			rpc, ok := status.Details()[0].(*rpc.Status)
			if !ok {
				t.Error(err)
				return
			}
			t.Error(rpc.Code, rpc.Message)
			return
		} else {
			t.Error(err)
			return
		}
	}
	t.Logf("获取 %d 条动态", len(dynAllResp.DynamicList.List))
	jsonString, err := json.Marshal(dynAllResp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(jsonString))
	return
}

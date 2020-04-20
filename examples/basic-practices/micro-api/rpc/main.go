package main

import (
	"context"
	"log"

	pb "examples/basic-practices/micro-api/rpc/proto" 	//⚠️ 导入路径
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
)

type Demo struct{}

// Call 方法会接收由API层转发，路由为/example/call的HTTP请求
func (e *Demo) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Print("收到请求")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.demo", "no content")
	}

	rsp.Message = "RPC响应 " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.rpc.demo"), //prefix.module
	)

	service.Init()

	// 注册 example 接口
	pb.RegisterDemoHandler(service.Server(), new(Demo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//1. 根据proto协议生成接口代码
// 命令：protoc --go_out=. --micro_out=. ./proto/*.proto
// go_out生成结构体对象文件
// micro_out生成接口服务方法等

// 2. 启动Micro服务：
// 命令：go run main.go

// 3.Micro调用
// 格式：micro prefix.module struct.method <jsonStr>
// 命令：micro call go.micro.rpc.demo Demo.Call '{"name":"史布斯"}'

// 4. 启动Api网关
// 格式：micro api --handler=rpc --namespace=prefix
// 命令：micro api -handler=rpc --namespace=go.micro.rpc

// 5. Http调用
// Http路径格式：module/struct/method
//curl -H 'Content-Type: application/json' -d '{"name": "史布斯"}' "http://localhost:8080/demo/demo/call"
//curl -H 'Content-Type: application/json' -d '{"name": "史布斯"}' "http://localhost:8080/demo/call"
// 说明：module名称和struct名称相同是可省略一个名称


//总结：
// RPC网关下API只接收POST方式的请求；
// Rpc的http内容格式content-type为application/json或者application/protobuf；
// Rpc网关最终是转发到了rpc服务上了；


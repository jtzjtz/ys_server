package main

import (
	"flag"
	"fmt"
	"github.com/jtzjtz/ys_pack/proto/department_proto"
	"github.com/jtzjtz/ys_pack/proto/sq_user_blacklist_proto"
	"github.com/jtzjtz/ys_pack/server/controller/department_controller"
	"github.com/jtzjtz/ys_pack/server/controller/sq_user_blacklist_controller"
	"github.com/jtzjtz/ys_server/app/config"
	"github.com/jtzjtz/ys_server/app/sqldb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var env = flag.String("env", "", "请输入运行环境:\n dev:测试分支\n fat:测试主干\n uat:预上线环境\n prod-aliyun:正式环境\n")
var server = flag.String("server", "rpc", "rpc or http\n")
var rpcAddr = flag.String("rpc_addr", ":8090", "rpc port\n")

func init() {
	flag.Parse()

	//// 验证
	//if *env == "" {
	//	fmt.Println("env启动参数异常！")
	//	//os.Exit(1)
	//}

	// 初始化 配置
	config.Init(*env)

	//	// 初始化 MySQL 连接池
	sqldb.Init()
	//
	//	// 初始化 Redis 连接池
	//	redisdb.Init()
}

func main() {
	// 启动 gRPC 服务
	fmt.Println("gRPC Server Listen on " + *rpcAddr)

	listen, err := net.Listen("tcp", *rpcAddr)
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
	}

	// 实例化 gRPC Server
	s := grpc.NewServer()

	// 心跳检测
	healthpb.RegisterHealthServer(s, health.NewServer())

	// 服务注册
	department_proto.RegisterDepartmentServiceServer(s, &department_controller.DepartmentController{})
	sq_user_blacklist_proto.RegisterSqUserBlacklistServiceServer(s, &sq_user_blacklist_controller.SqUserBlacklistController{})
	//sq_admin_menu_proto.RegisterSqAdminMenuServiceServer(s, &sq_admin_menu_controller.SqAdminMenuController{})
	//sq_user_blacklist_proto.RegisterSqUserBlacklistServiceServer(s, &sq_user_blacklist_controller.SqUserBlacklistController{})
	//sq_ugc_log_proto.RegisterSqUgcLogServiceServer(s, &sq_ugc_log_controller.SqUgcLogController{})
	// 服务反射（非必须）
	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to gRPC server: %v", err)
	}

}

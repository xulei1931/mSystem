package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	dbconfig "mSystem/src/common/db"
	"mSystem/src/common/pb"
	register "mSystem/src/user/compone"
	edpts "mSystem/src/user/endpoint"
	"mSystem/src/user/service"
	transport "mSystem/src/user/transport"
	"mSystem/src/utils"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		consulHost  = flag.String("consul.host", "127.0.0.1", "consul ip address")
		consulPort  = flag.String("consul.port", "8500", "consul port")

		serviceHost = flag.String("service.host", "localhost", "service ip address")
		servicePort = flag.String("service.port", "9001", "service port")

		grpcAddr    = flag.String("grpc", ":8001", "gRPC listen address.")
	)
	flag.Parse()
	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// 接口定义
	// 具体服务实现了
	svc := service.NewUserService(utils.GetLogger())
	//创建Endpoint
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)

	endpoint := edpts.NewEndpoint(svc,utils.GetLogger(),golangLimit)
	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpoint, logger)

	//创建注册对象
	registar := register.Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

    // http 服务
	go func() {
		fmt.Println("Http Server start at port:" + *servicePort)
		//启动前执行注册
		registar.Register()

		errChan <- http.ListenAndServe(":"+*servicePort, r)
	}()
	// 数据库连接初始化。。。。。。
	dbconfig.Init()
	//grpc server
	go func() {
		fmt.Println("grpc Server start at port" + *grpcAddr)
		listener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}
		baseServer := grpc.NewServer()
		pb.RegisterUserServiceExtServer(baseServer, svc)
		errChan <- baseServer.Serve(listener)

	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	registar.Deregister()
	fmt.Println(error)
}
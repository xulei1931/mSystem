package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	dbconfig "mSystem/svc/common/db"
	"mSystem/svc/common/pb"
	register "mSystem/svc/movie/compone"
	edpts "mSystem/svc/movie/endpoint"
	"mSystem/svc/movie/service"
	transport "mSystem/svc/movie/transport"
	"mSystem/svc/utils"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		consulHost = flag.String("consul.host", "127.0.0.1", "consul ip address")
		consulPort = flag.String("consul.port", "8500", "consul port")

		serviceHost = flag.String("service.host", "localhost", "service ip address")
		servicePort = flag.String("service.port", "9002", "service port")

		grpcAddr = flag.String("grpc", ":8002", "gRPC listen address.")

		zipkinURL = flag.String("zipkin.url", "http://127.0.0.1:9411/api/v2/spans", "Zipkin server url")
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
	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = *serviceHost + ":" + *servicePort
			serviceName   = "user-service"
			useNoopTracer = (*zipkinURL == "")
			reporter      = zipkinhttp.NewReporter(*zipkinURL)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracer, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", *zipkinURL)
		}
	}

	// 接口定义
	// 具体服务实现了
	svc := service.NewMovieService(utils.GetLogger())
	//创建Endpoint
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)
	// 创建endpoint
	endpoint := edpts.NewEndpoint(svc, utils.GetLogger(), golangLimit, zipkinTracer)

	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpoint, logger, zipkinTracer)

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
		pb.RegisterFilmServiceExtServer(baseServer, svc)
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
	dbconfig.Close() // 关闭数据库资源。。。。。。。
	fmt.Println(error)
}

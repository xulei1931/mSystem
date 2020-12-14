package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/openzipkin/zipkin-go"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	pb "mSystem/svc/common/pb"
	"mSystem/svc/file/service"
	"mSystem/svc/file/encode"
)
type HotPlayMoviesRequest struct {
	Location    int64 `json:"location"`
	MovieId int64 `json:"movieId"`

}
type MovieDetailRequest struct{
	MovieId int64 `json:"movieId"`
}
type MovieCreditsWithTypesRequest struct {
	MovieId int64 `json:"movieId"`
}

// y有几个函数就定义几个 endpoint
type Endpoints struct {
	HotPlayMovies      endpoint.Endpoint
	MovieDetail       endpoint.Endpoint
	MovieCreditsWithTypes endpoint.Endpoint
}
func NewEndpoint(s service.FileService, log *zap.Logger, limit *rate.Limiter,tracer *zipkin.Tracer) Endpoints {
	var HotPlayMoviesEndPoint endpoint.Endpoint
	HotPlayMoviesEndPoint = MakeHotPlayMoviesEndPointEndPoint(s)
	HotPlayMoviesEndPoint = kitzipkin.TraceEndpoint(tracer, "register-endpoint")(HotPlayMoviesEndPoint) //// 链路追踪

	var MovieDetailPoint endpoint.Endpoint
	MovieDetailPoint = MakeMovieDetailPoint(s)
//	LoginEndPoint = LoggingMiddleware(log)(LoginEndPoint)              // 登陆中间价
//	LoginEndPoint = NewGolangRateAllowMiddleware(limit)(LoginEndPoint) //限流
//	LoginEndPoint = kitzipkin.TraceEndpoint(tracer, "login-endpoint")(LoginEndPoint) //// 链路追踪

	var MovieCreditsWithTypesPoint endpoint.Endpoint
	MovieCreditsWithTypesPoint = MakeMovieCreditsWithTypesEndPoint(s)
	MovieCreditsWithTypesPoint = kitzipkin.TraceEndpoint(tracer, "getuserinfo-endpoint")(MovieCreditsWithTypesPoint) //// 链路追踪

	return Endpoints{HotPlayMovies: HotPlayMoviesEndPoint, MovieDetail: MovieDetailPoint, MovieCreditsWithTypes: MovieCreditsWithTypesPoint}
}
func MakeHotPlayMoviesEndPointEndPoint(s service.FileService ) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*HotPlayMoviesRequest) ///// 获取请求参数

		val, err := s.HotPlayMovies(ctx, &pb.HotPlayMoviesReq{ // 组装请求参数到servce
			Location:    req.Location,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}
func MakeMovieDetailPoint(s service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*MovieDetailRequest) ///// 获取请求参数

		val, err := s.MovieDetail(ctx, &pb.MovieDetailReq{
			MovieId: req.MovieId,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}
func MakeMovieCreditsWithTypesEndPoint(s service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*MovieCreditsWithTypesRequest) ///// 获取请求参数

		val, err := s.MovieCreditsWithTypes(ctx, &pb.MovieCreditsWithTypesReq{
			MovieId: req.MovieId,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}

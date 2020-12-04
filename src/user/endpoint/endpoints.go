package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	pb "mSystem/src/common/pb"
	"mSystem/src/user/encode"
	"mSystem/src/user/service"
)

type RegistRequest struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type RegistResponse struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetTokenRequest struct {
	Token string `json:"token"`
}

// y有几个函数就定义几个 endpoint
type Endpoints struct {
	RegistAccount      endpoint.Endpoint
	LoginAccount       endpoint.Endpoint
	GetUserInfoByToken endpoint.Endpoint
}

func NewEndpoint(s service.UserService, log *zap.Logger, limit *rate.Limiter) Endpoints {
	var RegistEndPoint endpoint.Endpoint
	RegistEndPoint = MakeRegistEndPoint(s)

	var LoginEndPoint endpoint.Endpoint
	LoginEndPoint = MakeLoginEndPoint(s)
	LoginEndPoint = LoggingMiddleware(log)(LoginEndPoint)              // 登陆中间价
	LoginEndPoint = NewGolangRateAllowMiddleware(limit)(LoginEndPoint) //限流

	var GetUserInfoByToken endpoint.Endpoint
	GetUserInfoByToken = MakeTokenEndPoint(s)
	return Endpoints{RegistAccount: RegistEndPoint, LoginAccount: LoginEndPoint, GetUserInfoByToken: GetUserInfoByToken}
}

// 实现请求转发
func MakeRegistEndPoint(s service.UserService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RegistRequest) ///// 获取请求参数

		val, err := s.RegistAccount(ctx, &pb.RegistAccountReq{ // 组装请求参数到servce
			Email:    req.Email,
			Password: req.Password,
			UserName: req.UserName,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}

// 实现请求转发
func MakeLoginEndPoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest) /// 获取请求参数
		val, err := s.LoginAccount(ctx, &pb.LoginAccountReq{ //组装请求参数到servce
			Email:    req.Email,
			Password: req.Password,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}
func MakeTokenEndPoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetTokenRequest) /// 获取请求参数
		val, err := s.GetUserInfoByToken(ctx, &pb.GetUserInfoByTokenRequest{
			Token: req.Token,
		})
		return encode.Response{
			Error: err,
			Data:  val,
		}, err
	}
}

package service

import (
	"context"
	"go.uber.org/zap"
	"mSystem/src/common/pb"
	"fmt"
)

const ContextReqUUid = "req_uuid"

type NewMiddlewareServer func(UserService) UserService

type logMiddlewareServer struct {
	logger *zap.Logger
	next   UserService
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service UserService) UserService {
		return &logMiddlewareServer{
			logger: log,
			next: service,
		}
	}
}

func (l logMiddlewareServer) RegistAccount(ctx context.Context, in *pb.RegistAccountReq) (out *pb.RegistAccountRsp,  err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint("请求uid：",ctx.Value(ContextReqUUid)), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("err", err))
	}()
	out, err = l.next.RegistAccount(ctx, in)
	return
}

func (l logMiddlewareServer) LoginAccount(ctx context.Context, in *pb.LoginAccountReq) (out *pb.LoginAccountRsp,  err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint("请求uid：",ctx.Value(ContextReqUUid)), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("err", err))
	}()
	out, err =  l.next.LoginAccount(ctx, in)
	return
}
func (l logMiddlewareServer) GetUserInfoByToken(ctx context.Context, in *pb.GetUserInfoByTokenRequest) (out *pb.GetUserInfoByTokenResponse,  err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint("请求uid：",ctx.Value(ContextReqUUid)), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("err", err))
	}()
	out, err =  l.next.GetUserInfoByToken(ctx, in)
	return
}
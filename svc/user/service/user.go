package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mSystem/svc/common/db"
	"mSystem/svc/common/errors"
	"mSystem/svc/common/pb"
	"mSystem/svc/utils"
)

type UserService interface {
	RegistAccount(ctx context.Context, req *pb.RegistAccountReq) (*pb.RegistAccountRsp, error)
	LoginAccount(ctx context.Context, req *pb.LoginAccountReq) (*pb.LoginAccountRsp, error)
	GetUserInfoByToken(ctx context.Context, req *pb.GetUserInfoByTokenRequest) (*pb.GetUserInfoByTokenResponse, error)
}
type baseServer struct {
	logger *zap.Logger
}

func NewUserService(log *zap.Logger) UserService {
	var server UserService
	server = &baseServer{
		logger: log,
	}
	//server = NewLogMiddlewareServer(log)(server)
	return server
}

// 账户注册
func (u baseServer) RegistAccount(ctx context.Context, req *pb.RegistAccountReq) (*pb.RegistAccountRsp, error) {
	userName := req.UserName
	password := utils.Md5Password(req.Password) // 密码md5加密
	email := req.Email
	user, err := db.SelectUserByEmail(email)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return &pb.RegistAccountRsp{
			Code: -1,
		}, errors.ErrorUserFailed
	}
	if user != nil {
		return &pb.RegistAccountRsp{
			Code: -1,
		}, errors.ErrorUserAlready
	}
	err = db.InsertUser(userName, password, email)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return &pb.RegistAccountRsp{
			Code: -1,
		}, errors.ErrorUserFailed
	}
	return &pb.RegistAccountRsp{
		Code: 0,
	}, nil
}

func (u baseServer) LoginAccount(ctx context.Context, req *pb.LoginAccountReq) (*pb.LoginAccountRsp, error) {
	email := req.Email
	password :=  utils.Md5Password(req.Password)
	user, err := db.SelectUserByPasswordName(email, password)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return &pb.LoginAccountRsp{
		}, errors.ErrorUserFailed
	}
	if user == nil {
		return &pb.LoginAccountRsp{}, errors.ErrorUserLoginFailed
	}
	// jwt 加密
	Token, err := utils.CreateJwtToken(user.UserName, int(user.UserId))
	return &pb.LoginAccountRsp{
		Token: Token,
		Code:  0,
		Uid:   fmt.Sprint(ctx.Value(ContextReqUUid)),
	}, err
}

func (u baseServer) GetUserInfoByToken(ctx context.Context, req *pb.GetUserInfoByTokenRequest) (*pb.GetUserInfoByTokenResponse, error) {
	token := req.Token
	if token == "" {
		return &pb.GetUserInfoByTokenResponse{
			Code: -1,
		}, errors.ErrorTokenEmpty
	}
	MapClaims, err := utils.ParseToken(token)
	fmt.Println("MapClaims",  MapClaims, MapClaims["DcId"], MapClaims["Name"],err)
	if err != nil {
		return &pb.GetUserInfoByTokenResponse{
			Code: -2,
		}, errors.FormatError("user",err.Error())
	}
	user_id := MapClaims["DcId"].(float64)
	username := MapClaims["Name"].(string)
	fmt.Println("user-info", user_id, username)
	user_info, e := db.SelectUserById(user_id, username)
	if e != nil {
		return &pb.GetUserInfoByTokenResponse{
			Code: -2,
		}, errors.ErrorUserLoginFailed
	}
	// redis
	return &pb.GetUserInfoByTokenResponse{
		Code: 0,
		UserInfo: &pb.UserInfo{
			UserId:   user_info.UserId,
			UserName: user_info.UserName,
			Email:    user_info.Email,
			Phone:    user_info.Phone,
		},
	}, nil
}

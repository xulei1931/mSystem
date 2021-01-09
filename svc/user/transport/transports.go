package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"user/encode"
	"user/endpoint"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

const ContextReqUUid = "req_uuid"

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints endpoint.Endpoints, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	// 链路追踪
//	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer, zipkin.Name("http-transport"))
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
		kithttp.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			logger.Log(fmt.Sprint(ctx.Value(ContextReqUUid)))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(err)
		}),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4(),nil), "req_uuid").String()
			logger.Log("给请求添加uuid", zap.Any("UUID", UUID))
			ctx = context.WithValue(ctx, ContextReqUUid, UUID)
			return ctx
		}),
	//	zipkinServer,
	}

     // 暴露具体的 endpoint
	r.Methods("POST").Path("/register").Handler(kithttp.NewServer(
		endpoints.RegistAccount,
		decodeRegisterRequest, // 请求参数
		encode.JsonResponse,
		options...,
	))

	r.Methods("POST").Path("/login").Handler(kithttp.NewServer(
		endpoints.LoginAccount,
		decodeLoginRequest, // 请求参数
		encode.JsonResponse,
		options...,
	))
	r.Methods("POST").Path("/userInfo").Handler(kithttp.NewServer(
		endpoints.GetUserInfoByToken,
		decodeGetTokenRequest, // 请求参数
		encode.JsonResponse,
		options...,
	))
	return r
}

// decodeStringRequest decode request params to struct
func decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json:", string(body))

	var rhe endpoint.RegistRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return &endpoint.RegistRequest{
		Email:    rhe.Email,
		UserName: rhe.UserName,
		Password: rhe.Password,
	}, nil
}
// 注册 decodeStringRequest decode request params to struct
func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json:", string(body))

	var rhe endpoint.LoginRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return &endpoint.LoginRequest{
		Email:    rhe.Email,
		Password: rhe.Password,
	}, nil
}
func decodeGetTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json:", string(body))

	var rhe endpoint.GetTokenRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return &endpoint.GetTokenRequest{
		Token:    rhe.Token,
	}, nil
}
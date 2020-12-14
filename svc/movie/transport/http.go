package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	gozipkin "github.com/openzipkin/zipkin-go"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"mSystem/svc/movie/encode"
	"mSystem/svc/movie/endpoint"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

const ContextReqUUid = "req_uuid"

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints endpoint.Endpoints, logger log.Logger, zipkinTracer *gozipkin.Tracer,) http.Handler {
	r := mux.NewRouter()
	// 链路追踪
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer, zipkin.Name("http-transport"))
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
		zipkinServer,
	}

     // 暴露具体的 endpoint
	r.Methods("POST").Path("/hot-movie").Handler(kithttp.NewServer(
		endpoints.HotPlayMovies,
		decodeHotPlayMoviesrRequest, // 请求参数
		encode.JsonResponse,
		options...,
	))

	r.Methods("POST").Path("/movie-detail").Handler(kithttp.NewServer(
		endpoints.MovieDetail,
		decodMovieDetailRequest, // 请求参数
		encode.JsonResponse,
		options...,
	))
	r.Methods("POST").Path("/credits").Handler(kithttp.NewServer(
		endpoints.MovieCreditsWithTypes,
		decodeMovieCreditsWithTypes, // 请求参数
		encode.JsonResponse,
		options...,
	))
	return r
}

// decodeStringRequest decode request params to struct
func decodeHotPlayMoviesrRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json-request:", string(body))

	var rhe endpoint.HotPlayMoviesRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	fmt.Println("request", rhe)

	return &rhe, nil
}
// 注册 decodeStringRequest decode request params to struct
func decodMovieDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json-request:", string(body))

	var rhe endpoint.MovieDetailRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return &rhe, nil
}
func decodeMovieCreditsWithTypes(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	println("json:", string(body))

	var rhe endpoint.MovieCreditsWithTypesRequest
	if err = json.Unmarshal(body, &rhe); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return &rhe, nil
}
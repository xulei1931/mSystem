package main

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"mSystem/svc/common/pb"
	"mSystem/svc/string-service/endpoint"
)

type grpcServer struct {
	diff grpc.Handler
}

func (s *grpcServer) Diff(ctx context.Context, r *pb.StringRequest) (*pb.StringResponse, error) {
	_, resp, err := s.diff.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.StringResponse), nil

}
func DecodeGRPCStringRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.StringRequest)
	return endpoint.StringRequest{
		RequestType: string(req.RequestType),
		A:           string(req.A),
		B:           string(req.B),
	}, nil
}
func EncodeGRPCStringResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.StringResponse)

	if resp.Error != nil {
		return &pb.StringResponse{
			Result: string(resp.Result),
			Err:    resp.Error.Error(),
		}, nil
	}

	return &pb.StringResponse{
		Result: string(resp.Result),
		Err:    "",
	}, nil
}

func NewGRPCServer(ctx context.Context, endpoints endpoint.StringEndpoints, serverTracer grpc.ServerOption) pb.StringServiceServer {
	return &grpcServer{
		diff: grpc.NewServer(
			endpoints.StringEndpoint,
			DecodeGRPCStringRequest,
			EncodeGRPCStringResponse,
			serverTracer,
		),
	}
}

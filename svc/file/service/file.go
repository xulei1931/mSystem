package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mSystem/svc/common/db"
	"mSystem/svc/common/pb"
)

type FileService interface {
	//获取正在售票的影片
	HotPlayMovies(ctx context.Context, req *pb.HotPlayMoviesReq) (*pb.HotPlayMoviesRep, error)
	// 电影详情
	MovieDetail(ctx context.Context, req *pb.MovieDetailReq) (*pb.MovieDetailRep, error)
	//MovieCreditsWithTypes
	MovieCreditsWithTypes(ctx context.Context, req *pb.MovieCreditsWithTypesReq) (*pb.MovieCreditsWithTypesRep, error)
}
type baseServer struct {
	logger *zap.Logger
}

func NewFileService(log *zap.Logger) baseServer {
	return baseServer{
		logger: log,
	}
}

func (a baseServer) HotPlayMovies(ctx context.Context, req *pb.HotPlayMoviesReq) (*pb.HotPlayMoviesRep, error) {

	return &pb.HotPlayMoviesRep{
		Movies: nil,
	}, nil
}

func (a baseServer) MovieDetail(ctx context.Context, req *pb.MovieDetailReq) (*pb.MovieDetailRep, error) {
	movieId := req.MovieId
	file, error := db.SelectFilmDetail(movieId)
	if error != nil {
		fmt.Println(file,error,"444",error.Error())

		return &pb.MovieDetailRep{
			Code: -1,
			Err:  string(error.Error()), // ###转为字符串
		}, nil
	}
	return &pb.MovieDetailRep{
		Code: 0,
		Res: &pb.MovieDetailInfo{
			TitleCn: file.TitleCn,
			Image:   file.Img,
			TitleEn: file.TitleEn,
			Rating:  file.RatingFinal,
			Year:    file.RYear,
		},
		Err: "",
	}, nil
}
func (a baseServer) MovieCreditsWithTypes(ctx context.Context, req *pb.MovieCreditsWithTypesReq) (*pb.MovieCreditsWithTypesRep, error) {
	return nil, nil
}

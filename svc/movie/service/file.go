package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mSystem/svc/common/config"
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

func NewMovieService(log *zap.Logger) baseServer {
	return baseServer{
		logger: log,
	}
}
// 获取热播电影

func (a baseServer) HotPlayMovies(ctx context.Context, req *pb.HotPlayMoviesReq) (*pb.HotPlayMoviesRep, error) {

	files, error := db.SelectTickingFilims(config.TickingNow)
	if error != nil {
		fmt.Println(files, error, "444", error.Error())
	}
	movie := make([]*pb.Movie,len(files))
	if len(files)>0{
      for index,item := range files{
		  filmActors := []string{"荒坡","范冰冰"}
		  actors :=""
		  for _, filmActor := range filmActors {
			  actors = actors + filmActor + " "
		  }

		  movie[index]= &pb.Movie{
			  Actors:actors,
             TitleCn: item.TitleCn,
		  }
	  }
	}
	fmt.Println(files, error, "555", error)
	return &pb.HotPlayMoviesRep{
		Movies:movie ,
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

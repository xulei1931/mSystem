package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"movie/common/db"
	"movie/common/pb"
)

type FileService interface {
	// 电影分类
	GetMovieTags(ctx context.Context, req *pb.MoviesTagsReq) (*pb.MoviesTagsRep, error)
	//获取正在售票的影片
	GetMoviesList(ctx context.Context, req *pb.MoviesListReq) (*pb.MoviesListsRep, error)
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
func (a baseServer) GetMovieTags(ctx context.Context, req *pb.MoviesTagsReq) (*pb.MoviesTagsRep, error) {
	tags, _, error := db.GetTags(int64(req.TagId))
	if error !=nil{
		return &pb.MoviesTagsRep{
         Code: -1,
         Err: error.Error(),
		},nil
	}
	tagsData := make([]*pb.TagsInfo, len(tags))

	for index,tag := range tags{
		tagsData[index] = &pb.TagsInfo{
           TagId:int32(tag.Id),
           CateName: tag.CatName,
           Sort: int32(tag.Sort),
		}
	}
	return &pb.MoviesTagsRep{
		Code: 0,
		List: tagsData,
	},nil
}
// 获取热播电影

func (a baseServer) GetMoviesList(ctx context.Context, req *pb.MoviesListReq) (*pb.MoviesListsRep, error) {

	files, count, error := db.MovieLists(req.TagId, req.Limit, req.Page)
	if error != nil {
		fmt.Println(files, error, "444", error.Error())
	}
	movie := make([]*pb.MovieDetailInfo, len(files))
	if len(files) > 0 {
		for index, item := range files {
			var PlayAble int64
			var IsNew int64
			if item.PlayAble {
				PlayAble = 1
			}
			if item.IsNew {
				IsNew = 1
			}
			movie[index] = &pb.MovieDetailInfo{
				MovieId:  item.MovieId,
				TagId:    int64(item.TagId),
				Title:    item.Title,
				Url:      item.Url,
				Cover:    item.Cover,
				Rate:     item.Rate,
				PlayAble: PlayAble,
				IsNew:    IsNew,
			}
		}
	}
	fmt.Println(req.TagId, files, count, error, "555", error)
	return &pb.MoviesListsRep{
		List: movie,
		Count:  count,
		Code:   0,
	}, nil
}

func (a baseServer) MovieDetail(ctx context.Context, req *pb.MovieDetailReq) (*pb.MovieDetailRep, error) {
	movieId := req.MovieId
	file, error := db.SelectFilmDetail(movieId)
	if error != nil {
		fmt.Println(file, error, "444", error.Error())

		return &pb.MovieDetailRep{
			Code: -1,
			Err:  string(error.Error()), // ###转为字符串
		}, nil
	}
	var (
		PlayAble int64
		IsNew int64
	)
	if file.PlayAble {
		PlayAble = 1
	}
	if file.IsNew {
		IsNew = 1
	}
	return &pb.MovieDetailRep{
		Code: 0,
		Res: &pb.MovieDetailInfo{
			MovieId:  file.MovieId,
			TagId:    int64(file.TagId),
			Title:    file.Title,
			Url:      file.Url,
			Cover:    file.Cover,
			Rate:     file.Rate,
			PlayAble: PlayAble,
			IsNew:    IsNew,
		},
		Err: "",
	}, nil
}
func (a baseServer) MovieCreditsWithTypes(ctx context.Context, req *pb.MovieCreditsWithTypesReq) (*pb.MovieCreditsWithTypesRep, error) {
	return nil, nil
}

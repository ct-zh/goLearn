package repositories

import "go_iris/datamodels"

// 定义 movie 的 方法
type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

// 实现 MovieRepository 的方法
func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name: "mukewang"}
	return movie.Name
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}


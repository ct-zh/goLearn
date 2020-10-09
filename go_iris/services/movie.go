package services

import (
	"go_iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManage struct {
	repo repositories.MovieRepository
}

func (m *MovieServiceManage) ShowMovieName() string {
	return m.repo.GetMovieName()
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManage{repo: repo}
}

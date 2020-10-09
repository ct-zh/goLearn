package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go_iris/repositories"
	"go_iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	return mvc.View{
		Name: "movie/index.html",
		Data: movieService.ShowMovieName(),
	}
}

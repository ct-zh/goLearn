package main

import (
	"log"
	"net/http"
	"web/web"
)

func main() {
	engine := web.New()

	engine.Get("/", func(c *web.Context) {
		c.Html(http.StatusOK, "<h1>Hello world</h1>")
	})

	engine.Get("/hello", func(c *web.Context) {
		c.String(http.StatusOK, "hello %s you're at %s\n", c.Query("name"), c.Path)
	})

	engine.Post("/login", func(ctx *web.Context) {
		ctx.Json(http.StatusOK, web.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	log.Fatal(engine.Run(":9999"))
}

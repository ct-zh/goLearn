package demo

import (
	"log"
	"net/http"
)

// 一个中间件demo

type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{mux: make(map[string]http.Handler)}
}

func (r *Router) Use(m middleware) {
	log.Println("add middleware ")
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(router string, h http.Handler) {
	var mergeHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}
	r.mux[router] = mergeHandler

	log.Println("add router: ", router)
}

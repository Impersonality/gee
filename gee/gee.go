package gee

import "net/http"

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.router.addRouter("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.router.addRouter("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

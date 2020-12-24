package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:      w,
		R:      req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) SetStatus(code int) {
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

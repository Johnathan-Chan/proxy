package core

import (
	"net/http"
)

type Context struct {
	Request *http.Request
	Writer http.ResponseWriter
	BeforeRequest HandlersChain
	indexBefore int
}

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

func (c *Context) Reset(w http.ResponseWriter, r *http.Request){
	c.Request = r
	c.Writer = w
	c.BeforeRequest = nil
	c.indexBefore = -1
}

func (c *Context) NextBefore()  {
	c.indexBefore++
	for c.indexBefore < len(c.BeforeRequest) {
		c.BeforeRequest[c.indexBefore](c)
		c.indexBefore++
	}
}

func (c *Context) Abort()  {
	c.indexBefore = len(c.BeforeRequest)
}

func (c *Context) Copy() *Context {
	cp := Context{
		Writer: c.Writer,
		Request:   c.Request,
	}

	return &cp
}
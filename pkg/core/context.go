package core

import (
	"net/http"
)

type Context struct {
	Request *http.Request
	ResponseWriter http.ResponseWriter
	BeforeRequest HandlersChain
	AfterResponse HandlersChain
	indexBefore int
	indexAfter int
}

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

func (c *Context) Reset(w http.ResponseWriter, r *http.Request){
	c.Request = r
	c.ResponseWriter = w
	c.BeforeRequest = nil
	c.AfterResponse = nil
	c.indexBefore = -1
	c.indexAfter = -1
}

func (c *Context) NextBefore()  {
	c.indexBefore++
	for c.indexBefore < len(c.BeforeRequest) {
		c.BeforeRequest[c.indexBefore](c)
		c.indexBefore++
	}
}

func (c *Context) NextAfter() {
	c.indexAfter++
	for c.indexAfter < len(c.AfterResponse) {
		c.AfterResponse[c.indexAfter](c)
		c.indexAfter++
	}
}

func (c *Context) Abort()  {
	c.indexBefore = len(c.BeforeRequest)
}


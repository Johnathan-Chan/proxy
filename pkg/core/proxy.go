package core

import (
	"io"
	"net/http"
	"sync"
)

type Proxy struct {
	pool sync.Pool
	beforeRequest HandlersChain
	afterResponse HandlersChain
}

func NewProxy() *Proxy{
	proxy := &Proxy{}
	proxy.pool.New = func() interface{} {
		return &Context{}
	}
	return proxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := p.pool.Get().(*Context)
	c.Reset(w, r)
	p.handler(c)
	p.pool.Put(c)
}

func (p *Proxy) handler(ctx *Context) {
	ctx.BeforeRequest = p.beforeRequest
	ctx.NextBefore()
}

func (p *Proxy) UseBefore(before ...HandlerFunc) {
	p.beforeRequest = append(p.beforeRequest, before...)
}


func (p *Proxy) Handler(core HandlerFunc){
	p.beforeRequest = append(p.beforeRequest, core)
}

func Transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func CopyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

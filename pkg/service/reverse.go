package service

import (
	"fmt"
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(ctx *core.Context) {
	urls := fmt.Sprintf("http://127.0.0.1:8081")
	remote, err := url.Parse(urls)
	if err != nil {
		log.Println("Proxy Error: ", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	tmp := proxy.Director
	proxy.Director = func(r *http.Request) {
		tmp(r)
		r.Host = remote.Host
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

package main

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"github.com/Johnathan-Chan/proxy/pkg/middlewares"
	"github.com/Johnathan-Chan/proxy/pkg/service"
	"net/http"
)

func main(){

	proxy := core.NewProxy()
	proxy.UseBefore(middlewares.Cors())
	proxy.Handler(service.ReverseProxy)

	server := &http.Server{
		Addr:      ":8080",
		Handler:   proxy,
	}

	if err := server.ListenAndServe(); err != nil{
		panic(err)
	}
}




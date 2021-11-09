package main

import (
	"crypto/tls"
	"github.com/Johnathan-Chan/proxy/pkg/conf"
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"github.com/Johnathan-Chan/proxy/pkg/middlewares"
	"github.com/Johnathan-Chan/proxy/pkg/service"
	"github.com/Johnathan-Chan/proxy/pkg/utils"
	"net/http"
)

func main(){
	config := &conf.Config{}
	if err := utils.LoadYamlConfig("./pkg/conf", "config.yaml", &config); err != nil{
		panic(err)
	}
	conf.GlobalConfig = config

	cert, err := utils.GenCertificate()
	if err != nil {
		panic(err)
	}

	proxy := core.NewProxy()
	proxy.UseBefore(middlewares.HttpHeaderAuth())
	proxy.Handler(service.HandleHttpAndHttps)

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler:   proxy,
	}

	if err = server.ListenAndServe(); err != nil{
		panic(err)
	}
}

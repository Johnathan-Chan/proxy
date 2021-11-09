package service

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"net"
	"net/http"
	"time"
)

func HandleHttps(ctx *core.Context){
	destConn, err := net.DialTimeout("tcp", ctx.Request.Host, 60*time.Second)
	if err != nil {
		http.Error(ctx.ResponseWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ctx.ResponseWriter.WriteHeader(http.StatusOK)

	hijacker, ok := ctx.ResponseWriter.(http.Hijacker)
	if !ok {
		http.Error(ctx.ResponseWriter, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(ctx.ResponseWriter, err.Error(), http.StatusServiceUnavailable)
	}

	go core.Transfer(destConn, clientConn)
	go core.Transfer(clientConn, destConn)
}


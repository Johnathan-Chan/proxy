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
		http.Error(ctx.Writer, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)

	hijacker, ok := ctx.Writer.(http.Hijacker)
	if !ok {
		http.Error(ctx.Writer, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusServiceUnavailable)
	}

	go core.Transfer(destConn, clientConn)
	go core.Transfer(clientConn, destConn)
}


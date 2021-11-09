package service

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"io"
	"net/http"
)

func HandleHttp(ctx *core.Context){
	resp, err := http.DefaultTransport.RoundTrip(ctx.Request)
	if err != nil {
		http.Error(ctx.ResponseWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}

	if resp.Body != nil{
		defer resp.Body.Close()
	}

	core.CopyHeader(ctx.ResponseWriter.Header(), resp.Header)
	ctx.ResponseWriter.WriteHeader(resp.StatusCode)
	io.Copy(ctx.ResponseWriter, resp.Body)
}

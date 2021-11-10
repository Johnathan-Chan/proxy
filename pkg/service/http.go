package service

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"io"
	"net/http"
)

func HandleHttp(ctx *core.Context){
	resp, err := http.DefaultTransport.RoundTrip(ctx.Request)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusServiceUnavailable)
		return
	}

	if resp.Body != nil{
		defer resp.Body.Close()
	}

	core.CopyHeader(ctx.Writer.Header(), resp.Header)
	ctx.Writer.WriteHeader(resp.StatusCode)
	io.Copy(ctx.Writer, resp.Body)
}

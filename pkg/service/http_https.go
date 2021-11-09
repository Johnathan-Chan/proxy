package service

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"net/http"
)

func HandleHttpAndHttps(ctx *core.Context){
	if ctx.Request.Method == http.MethodConnect {
		HandleHttps(ctx)
	} else {
		HandleHttp(ctx)
	}
}

package middlewares

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"log"
)

var cors = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
	"Access-Control-Allow-Headers": "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken",
	"Access-Control-Expose-Headers": "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
	"Access-Control-Max-Age":  "172800",
	"Access-Control-Allow-Credentials": "false",
	"content-type": "application/json",
}

func Cors() core.HandlerFunc{
	return func(ctx *core.Context) {
		for key, value := range cors {
			if ctx.Writer.Header().Get(key) == "" {
				ctx.Writer.Header().Set(key, value)
			}
		}

		for key, value := range ctx.Writer.Header(){
			log.Println(key, ":", value)
		}
	}
}


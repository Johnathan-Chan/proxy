package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/Johnathan-Chan/proxy/pkg/conf"
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"log"
	"net/http"
	"reflect"
	"strings"
	"unsafe"
)

func HttpHeaderAuth() core.HandlerFunc{
	return func(ctx *core.Context) {
		auth := ctx.Request.Header.Get("Proxy-Authorization")
		auth = strings.Replace(auth, "Basic ", "", 1)
		target := fmt.Sprintf("%s:%s", conf.GlobalConfig.User, conf.GlobalConfig.Password)
		tmp := *(*reflect.StringHeader)(unsafe.Pointer(&target))
		targetByte := *(*[]byte)(unsafe.Pointer(&tmp))
		target = base64.StdEncoding.EncodeToString(targetByte)

		if auth != target {
			ctx.Abort()
			log.Println("认证失败")

			ctx.ResponseWriter.WriteHeader(http.StatusOK)
			ctx.ResponseWriter.Write([]byte("认证失败"))
			return
		}

		ctx.Request.Header.Del("Proxy-Authorization")
		log.Println("认证成功")
	}
}

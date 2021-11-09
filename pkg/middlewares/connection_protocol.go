package middlewares

import (
	"github.com/Johnathan-Chan/proxy/pkg/core"
	"log"
)

func ConnectionProtocol() core.HandlerFunc{
	return func(ctx *core.Context) {
		log.Println("test")
	}
}

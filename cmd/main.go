package main

import (
	"https-proxy/pkg/service/https_proxy"
)

func main(){
	https_proxy.Serve(":8080")
}

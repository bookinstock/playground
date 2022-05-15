package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	// 开启pprof，监听请求
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
		os.Exit(1)
	}
	// end pprof

	// do sth
}

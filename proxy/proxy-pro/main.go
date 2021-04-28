package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func main() {
	router := httprouter.New()
	origin, _ := url.Parse("http://localhost:14082/")

	path := "/*catchall" // path
	reverseProxy := httputil.NewSingleHostReverseProxy(origin)
	// localhost:9001/hello
	reverseProxy.Director = func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host) // host req.Header.Add("X-Origin-Host", origin.Host) // host
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host

		wildcardIndex := strings.IndexAny(path, "*")
		proxyPath := singleJoiningSlash(origin.Path, req.URL.Path[wildcardIndex:])
		if strings.HasSuffix(proxyPath, "/") && len(proxyPath) > 1 {
			proxyPath = proxyPath[:len(proxyPath)-1]
		}
		fmt.Printf("\n\n======== proxyPath:%s\n\n", proxyPath)
		req.URL.Path = proxyPath

		// DumpRequest 以 HTTP/1.x线表示形式返回给定的请求。
		// 它只能被服务器用来调试客户端请求。返回的表示只是一个近似值;
		// 解析为 http.Request 时，初始请求的某些细节会丢失。特别是头字段名称的顺序和大小写会丢失。
		// 多值头中值的顺序保持不变。HTTP/2 请求以 HTTP/1.x 形式转储，而不是以其原始二进制表示。
		body, _ := httputil.DumpRequest(req, true)
		fmt.Printf("final request\n\n %+v \n\n", string(body))
	}

	router.Handle("POST", path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reverseProxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":9001", router))
}

// X-Forwarded-Host (XFH) 是一个事实上的标准首部，用来确定客户端发起的请求中使用  Host 指定的初始域名。
// 反向代理（如负载均衡服务器、CDN等）的域名或端口号可能会与处理请求的源头服务器有所不同，
// 在这种情况下，X-Forwarded-Host 可以用来确定哪一个域名是最初被用来访问的。

// final request
//
// GET /v1/user/info HTTP/1.1
// Host: localhost:9001
// Accept: */*
// User-Agent: curl/7.64.1
// X-Forwarded-Host: localhost:9001
//
//2021/04/28 15:09:56 http: proxy error: dial tcp [::1]:9000: connect: connection refused

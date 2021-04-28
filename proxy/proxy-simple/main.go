package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	origin, _ := url.Parse("http://service:9000/")

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host) // 上游
		req.Header.Add("X-Origin-Host", origin.Host) // 源
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe("proxy:9001", nil))
}

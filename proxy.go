package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/darkhelmet/env"
)

var (
	port = env.IntDefault("PORT", 5000)
)

func main() {
	proxyUrl, err := url.Parse(env.String("PROXY_URL"))
	if err != nil {
		log.Fatalf("failed parsing url: %s", err)
	}
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = proxyUrl.Scheme
			req.URL.Host = proxyUrl.Host
			req.URL.Path = req.URL.Path
			req.Host = proxyUrl.Host
		},
	}
	log.Printf("listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
}

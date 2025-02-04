package main

import (
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Potato Proxy: PORT is unset")
	}

	PROXY_HOST := os.Getenv("PROXY_HOST")
	if PROXY_HOST == "" {
		log.Fatal("Potato Proxy: PROXY_HOST is unset")
	}

	target := &url.URL{
		Scheme: "http",
		Host:   PROXY_HOST,
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetXForwarded()
			r.SetURL(target)

			// preserve the original host
			r.Out.Host = r.In.Host
		},
	}

	logger := slog.Default()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Method + " " + r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	logger.Info("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Potato Proxy: PORT is unset")
	}

	PROXY_URL := os.Getenv("PROXY_URL")
	if PROXY_URL == "" {
		log.Fatal("Potato Proxy: PROXY_URL is unset")
	}

	if !strings.HasPrefix(PROXY_URL, "http://") {
		PROXY_URL = "http://" + PROXY_URL
	}

	u, err := url.Parse(PROXY_URL)
	if err != nil {
		log.Fatal("Potato Proxy: Error parsing PROXY_URL")
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	logger := slog.Default()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Method + " " + r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Potato Proxy: Listening on port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}

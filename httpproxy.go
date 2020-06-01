package utils

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

// CheckEvnHTTPProxy func
func CheckEvnHTTPProxy() {
	proxy := os.Getenv("http_proxy")
	if len(proxy) == 0 {
		proxy = os.Getenv("https_proxy")
		if len(proxy) == 0 {
			proxy = os.Getenv("all_proxy")
			if len(proxy) == 0 {
				return
			}
		}
	}
	proxy = strings.ToLower(proxy)
	if strings.HasPrefix(proxy, "http://") == false {
		proxy = "http://" + proxy
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return
	}
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
}

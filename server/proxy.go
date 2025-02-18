package server

import (
	"SkyRP/config"
	"io"
	"net/http"
	"sync"
)

var nodeIdx int
var mu sync.Mutex

func getNextUpstream(data config.RootConfigSchema) string {
	mu.Lock()
	defer mu.Unlock()

	if len(data.Server.Upstreams) == 0 {
		return ""
	}

	nodeIdx = (nodeIdx + 1) % len(data.Server.Upstreams)
	return data.Server.Upstreams[nodeIdx].URL
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	data, err := config.Get_config_data()

	if err != nil {
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}
	
	targetUrl := getNextUpstream(data)

	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, targetUrl + r.RequestURI, r.Body)

	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)

	if err != nil {
		http.Error(w, "Failed to reach backend", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
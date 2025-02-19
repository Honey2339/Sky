package server

import (
	"SkyRP/config"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var nodeIdxMap = make(map[string]int)
var mu sync.Mutex

func groupUpstream(data config.RootConfigSchema) map[string][]string {
	grouped := make(map[string][]string)
	for _, upstream := range data.Server.Upstreams {
		parsedUrl, err := url.Parse(upstream.URL)
		if err != nil {
			continue
		}
		grouped[parsedUrl.Path] = append(grouped[parsedUrl.Path], upstream.URL)
	}

	return grouped
}

func getNextUpstream(data config.RootConfigSchema, path string) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	
	grouped := groupUpstream(data)

	servers, exists := grouped[path]
	if !exists || len(servers) == 0 {
		return "", fmt.Errorf("no upstream available for path: %s", path)
	}

	idx := nodeIdxMap[path]
	nodeIdxMap[path] = (idx + 1) % len(servers)

	return servers[idx], nil
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	data, err := config.Get_config_data()

	if err != nil {
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}
	
	targetUrl, err := getNextUpstream(data, r.URL.Path)
	if err != nil {
		http.Error(w, "No available upstreams", http.StatusServiceUnavailable)
		fmt.Println("Error getting upstream:", err)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, targetUrl, r.Body)

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
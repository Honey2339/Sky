package server

import (
	"SkyRP/config"
	"io"
	"net/http"
	"sync"
)

var nodeIdxMap = make(map[string]int)
var mu sync.Mutex

func groupUpstream(data config.RootConfigSchema) map[string][]string {
	grouped := make(map[string][]string)

	for _, upstream := range data.Server.Upstreams {
		grouped[upstream.URL] = append(grouped[upstream.URL], upstream.URL)
	}

	return grouped
}

func getNextUpstream(data config.RootConfigSchema, path string) string {
	mu.Lock()
	defer mu.Unlock()
	
	grouped := groupUpstream(data)
	println("the Path variable : ", path)
	for upstreamPath, servers := range grouped {
		println("the upstreamPath variable : ", upstreamPath)
		
		if path == upstreamPath {
			idx := nodeIdxMap[upstreamPath]
			nodeIdxMap[upstreamPath] = (idx + 1) % len(servers)

			return servers[idx]
		}
	}

	return ""
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	data, err := config.Get_config_data()

	if err != nil {
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}
	
	targetUrl := getNextUpstream(data, r.URL.Path)

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
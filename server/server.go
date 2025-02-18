package server

import (
	"SkyRP/config"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
)

func HttpServer() {
	data , err := config.Get_config_data()

	if err != nil {
		return
	}

	PORT := ":" + strconv.Itoa(data.Server.Listen)

	log.Infof("Sky is running on %s", PORT)
	http.HandleFunc("/", ProxyHandler)
	http.ListenAndServe(PORT, nil)
}
package server

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	if os.Getenv("OLLAMA_METRICS") == "1" {
		go http.ListenAndServe("127.0.0.1:2112", promhttp.Handler())
	}
}

package http

import (
	"net/http"
)

// ServerConfig is the configuration to provide to create an http.Server with NewServer()
type ServerConfig struct {
	ListenAddress string
	Handler http.Handler
}

// NewServer returns a http.Server
func NewServer(cfg ServerConfig) *http.Server {
	srv := &http.Server{
		Addr: cfg.ListenAddress,
		Handler: cfg.Handler,
	}

	return srv
}
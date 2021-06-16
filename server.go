package mqtt

import (
	"context"
	"github.com/rs/cors"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true}).Handler(handler),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}